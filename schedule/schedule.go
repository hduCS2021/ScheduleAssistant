package schedule

import (
	"context"
	"github.com/hduCS2021/ScheduleAssistant/schedule/trigger"
	"github.com/hduCS2021/ScheduleAssistant/student"
	"time"
)

type Schedule struct {
	Name     string
	StuSet   *student.StuSet
	Triggers []*trigger.Trigger
	do       func(schedule *Schedule)
}

func New(name string, stuSet *student.StuSet, do func(schedule *Schedule)) *Schedule {
	return &Schedule{
		Name:   name,
		do:     do,
		StuSet: stuSet,
	}
}

func (s *Schedule) AddTrigger(t ...*trigger.Trigger) *Schedule {
	s.Triggers = append(s.Triggers, t...)
	return s
}

func (s *Schedule) AddStudent(stu student.Student) *Schedule {
	s.StuSet.Add(stu)
	return s
}

func (s *Schedule) RemoveStudent(stu student.Student) *Schedule {
	s.StuSet.Remove(stu)
	return s
}

func (s *Schedule) isTriggered(t time.Time) bool {
	if len(s.Triggers) == 0 {
		return false
	}
	for _, v := range s.Triggers {
		switch v.IsTriggered(t) {
		case trigger.Accept:
			return true
		case trigger.Reject:
			return false
		case trigger.PassThrough:
		}
	}
	return true
}

func (s *Schedule) Run(ctx context.Context) *Schedule {
	go func() {
		time.Sleep(60*time.Second - time.Duration(time.Now().Second()))
		t := time.NewTicker(time.Minute)
		for {
			select {
			case t1 := <-t.C:
				if s.isTriggered(t1) {
					go s.do(s)
				}
			case <-ctx.Done():
				t.Stop()
				return
			}
		}
	}()
	return s
}
