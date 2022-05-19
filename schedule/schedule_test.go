package schedule

import (
	"context"
	"github.com/hduCS2021/ScheduleAssistant/schedule/date"
	"github.com/hduCS2021/ScheduleAssistant/schedule/trigger"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	schedule := New("test", func(s *Schedule) {
		t.Log("do something")
		t.Log(s.students)
	})
	schedule.AddStudent("bs")
	tr := trigger.New()
	m, err := date.NewMinute("0-59")
	if err != nil {
		t.Error(err)
	}
	tr.Add(m)
	schedule.AddTrigger(tr)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	schedule.Run(ctx)

	time.Sleep(time.Second * 150)
}
