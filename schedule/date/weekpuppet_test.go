package date

import (
	"testing"
	"time"
)

func TestWeekPuppet(t *testing.T) {
	//双周
	t1 := time.Date(2022, 5, 15, 0, 0, 0, 0, time.Local)
	//单周
	t2 := time.Date(2022, 5, 16, 0, 0, 0, 0, time.Local)
	p := NewWeekPuppet(2)
	if p.IsValid(t2) {
		t.Error("双周校验失败")
	}
	p = NewWeekPuppet(1)
	if p.IsValid(t1) {
		t.Error("单周校验失败")
	}
	p = NewWeekPuppet(0)
	if p.IsValid(t1) || p.IsValid(t2) {
		t.Error("无法触发校验失败")
	}
	p = NewWeekPuppet(3)
	if !p.IsValid(t1) || !p.IsValid(t2) {
		t.Error("都能触发校验失败")
	}
}
