package date

import (
	"testing"
	"time"
)

func TestYear(t *testing.T) {
	y := NewYear("2012-2016,2022")
	in := []int{2010, 2011, 2012, 2013, 2015, 2016, 2021, 2022, 2023}
	out := []bool{false, false, true, true, true, true, false, true, false}
	constructTime := func(year int) time.Time {
		return time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	}
	for i, v := range in {
		if y.isValid(constructTime(v)) != out[i] {
			t.Error(v, "年校验错误")
		}
	}
}
