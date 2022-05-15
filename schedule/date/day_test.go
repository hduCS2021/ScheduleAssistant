package date

import (
	"testing"
	"time"
)

func TestDay(t *testing.T) {
	m, err := NewDay("10-20,30")
	if err != nil {
		t.Error(err)
		return
	}
	d := time.Date(2000, 1, 1, 3, 1, 5, 6, time.Local)
	for i := 1; i < 10; i++ {
		if m.IsValid(d) {
			t.Error("日", d.Day(), "校验失败")
		}
		d = d.Add(time.Hour * 24)
	}
	for i := 10; i < 21; i++ {
		if !m.IsValid(d) {
			t.Error("日", d.Day(), "校验失败")
		}
		d = d.Add(time.Hour * 24)
	}
	for i := 21; i < 30; i++ {
		if m.IsValid(d) {
			t.Error("日", d.Day(), "校验失败")
		}
		d = d.Add(time.Hour * 24)
	}
	if !m.IsValid(d) {
		t.Error("日", d.Day(), "校验失败")
	}
}
