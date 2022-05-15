package date

import (
	"time"
)

type Minute uint64

func NewMinute(minute string) (Minute, error) {
	var m Minute
	err := parseString(&m, 0, 59, minute)
	return m, err
}

//IsValid check if a minute in range
//input should be in range 0-59
func (m Minute) IsValid(t time.Time) bool {
	return (m>>t.Minute())%2 == 1
}

func (m Minute) GetTimeUnit() time.Duration {
	return time.Minute
}
