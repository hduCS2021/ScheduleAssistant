package trigger

import (
	"github.com/hduCS2021/ScheduleAssistant/schedule/date"
	"testing"
	"time"
)

func TestTrigger(t *testing.T) {
	tr := New()
	min := date.NewMinute("1-10")
	tr.Add(min).SetPrior(1).SetAction(PassThrough)

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
