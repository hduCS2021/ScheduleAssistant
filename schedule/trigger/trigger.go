package trigger

import (
	"github.com/hduCS2021/ScheduleAssistant/schedule/date"
	"time"
)

type Trigger struct {
	dates date.List
	prior int
}

func New() *Trigger {
	return &Trigger{
		prior: 0,
	}
}

func (tr *Trigger) IsValid(t time.Time) bool {
	return tr.dates.Check(t)
}
