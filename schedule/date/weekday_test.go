package date

import (
	"testing"
	"time"
)

func TestWeekday(t *testing.T) {
	weekday, err := NewWeekday("2-4,7")
	if err != nil {
		t.Error(err)
		return
	}
	t0 := time.Date(2022, 5, 15, 0, 0, 0, 0, time.Local)
	result := []bool{false, true, true, true, false, false, true}
	for i := 1; i <= 7; i++ {
		if weekday.IsValid(t0.Add(time.Hour*24*time.Duration(i))) != result[i-1] {
			t.Error("周次错误", i)
		}
	}
}
