package date

import (
	"testing"
	"time"
)

func TestMonth(t *testing.T) {
	m := NewMonth("3-5,9")
	constructTime := func(month int) time.Time {
		return time.Date(2022, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	}
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := []bool{false, false, true, true, true, false, false, false, true, false}
	for i, v := range in {
		if m.IsValid(constructTime(v)) != out[i] {
			t.Error(i, "月校验错误")
		}
	}
}
