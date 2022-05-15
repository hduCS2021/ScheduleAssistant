package date

import (
	"time"
)

type Hour uint32

func NewHour(hour string) (Hour, error) {
	var h Hour
	err := parseString(&h, 0, 23, hour)
	return h, err
}

//IsValid check if Hour in range
//input should be in range [0,23]
func (m Hour) IsValid(t time.Time) bool {
	return (m>>t.Hour())%2 == 1
}

func (m Hour) GetTimeUnit() time.Duration {
	return time.Hour
}
