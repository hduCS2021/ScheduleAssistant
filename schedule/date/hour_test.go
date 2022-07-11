package date

import (
	"testing"
	"time"
)

func TestHour(t *testing.T) {
	m := NewHour("3-8,12")
	d := time.Date(2000, 1, 2, 0, 1, 5, 6, time.Local)

	for i := 0; i < 3; i++ {
		if m.IsValid(d) {
			t.Error("小时", d.Hour(), "校验失败")
		}
		d = d.Add(time.Hour)
	}
	for i := 3; i < 9; i++ {
		if !m.IsValid(d) {
			t.Error("小时", d.Hour(), "校验失败")
		}
		d = d.Add(time.Hour)
	}
	for i := 9; i < 12; i++ {
		if m.IsValid(d) {
			t.Error("小时", d.Hour(), "校验失败")
		}
		d = d.Add(time.Hour)
	}
	if !m.IsValid(d) {
		t.Error("小时", d.Hour(), "校验失败")
	}
}
