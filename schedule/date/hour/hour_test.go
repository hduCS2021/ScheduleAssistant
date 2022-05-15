package hour

import (
	"github.com/hduCS2021/ScheduleAssistant/schedule/date/minute"
	"testing"
)

func TestHour(t *testing.T) {
	m, err := minute.NewMinute("3-8,12")
	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 3; i++ {
		if m.IsValid(i) {
			t.Error("小时", i, "校验失败")
		}
	}
	for i := 3; i < 9; i++ {
		if !m.IsValid(i) {
			t.Error("小时", i, "校验失败")
		}
	}
	for i := 9; i < 12; i++ {
		if m.IsValid(i) {
			t.Error("小时", i, "校验失败")
		}
	}
	if !m.IsValid(12) {
		t.Error("小时", 30, "校验失败")
	}
}
