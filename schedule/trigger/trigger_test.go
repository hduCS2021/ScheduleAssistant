package trigger

import (
	date "github.com/hduCS2021/ScheduleAssistant/schedule/date/minute"
	"testing"
	"time"
)

func TestTrigger(t *testing.T) {
	tr := New()
	min, err := date.NewMinute("1-10")
	if err != nil {
		t.Error(err)
		return
	}
	tr.Add(min)
	tr.SetPrior(1)
	tr.SetAction(PassThrough)
	if tr.IsTriggered(
		time.Date(
			2000, 1, 2, 3, 4, 5, 6, time.Local)) != PassThrough {
		t.Error("triggered uncorrected")
	}
	if tr.IsTriggered(
		time.Date(
			2000, 1, 2, 3, 11, 5, 6, time.Local)) != Reject {
		t.Error("triggered uncorrected")
	}

}
