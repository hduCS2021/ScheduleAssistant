package date

import (
	"time"
)

type Day uint32

func NewDay(day string) (Day, error) {
	var d Day
	err := parseString(&d, 1, 31, day)
	return d, err
}

func (d Day) IsValid(t time.Time) bool {
	return (d>>(t.Day()-1))%2 == 1
}

func (d Day) GetTimeUnit() time.Duration {
	return time.Hour * 24
}
