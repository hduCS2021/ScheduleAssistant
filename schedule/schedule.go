package schedule

import (
	"github.com/hduCS2021/ScheduleAssistant/schedule/trigger"
)

type Schedule struct {
	name    string
	Student []*string
	Trigger []*trigger.Trigger
}
