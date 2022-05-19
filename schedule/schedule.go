package schedule

import (
	"context"
	"github.com/hduCS2021/ScheduleAssistant/schedule/trigger"
	"time"
)

type Schedule struct {
	name string
	//todo: student class
	students []string
	trigger  []*trigger.Trigger
	do       func(schedule *Schedule)
}

func New(name string, do func(schedule *Schedule)) *Schedule {
	return &Schedule{
		name: name,
		do:   do,
	}
}

func (s *Schedule) AddTrigger(t ...*trigger.Trigger) {
	s.trigger = append(s.trigger, t...)
}

func (s *Schedule) AddStudent(stu ...string) {
	s.students = append(s.students, stu...)
}

func (s *Schedule) isTriggered(t time.Time) bool {
	if len(s.trigger) == 0 {
		return false
	}
	for _, v := range s.trigger {
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

func (s *Schedule) Run(ctx context.Context) {
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
}
