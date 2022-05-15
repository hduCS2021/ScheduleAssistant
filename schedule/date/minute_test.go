package date

import (
	"testing"
	"time"
)

func TestMinute(t *testing.T) {
	m, err := NewMinute("10-20,30")
	if err != nil {
		t.Error(err)
		return
	}
	d := time.Date(2000, 1, 2, 3, 1, 5, 6, time.Local)
	for i := 1; i < 10; i++ {
		if m.IsValid(d) {
			t.Error("分钟", d.Minute(), "校验失败")
		}
		d = d.Add(time.Minute)
	}
	for i := 10; i < 21; i++ {
		if !m.IsValid(d) {
			t.Error("分钟", d.Minute(), "校验失败")
		}
		d = d.Add(time.Minute)
	}
	for i := 21; i < 30; i++ {
		if m.IsValid(d) {
			t.Error("分钟", d.Minute(), "校验失败")
		}
		d = d.Add(time.Minute)
	}
	if !m.IsValid(d) {
		t.Error("分钟", d.Minute(), "校验失败")
	}
}
