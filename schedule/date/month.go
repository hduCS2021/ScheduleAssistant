package date

import (
	"time"
)

type Month uint16

func NewMonth(month string) (Month, error) {
	var m Month
	err := parseString(&m, 1, 12, month)
	return m, err
}

func (d Month) IsValid(t time.Time) bool {
	return (d>>(t.Month()-1))%2 == 1
}

func (d Month) GetTimeUnit() time.Duration {
	return time.Hour * 24 * 30
}
