package date

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Minute uint64

func NewMinute(minute string) Minute {
	var m Minute
	if err := parseString(&m, 0, 59, minute); err != nil {
		log.Warnf("fail to parse time:%v", err)
	}
	return m
}

//IsValid check if a minute in range
//input should be in range 0-59
func (m Minute) IsValid(t time.Time) bool {
	return (m>>t.Minute())%2 == 1
}

func (m Minute) GetTimeUnit() time.Duration {
	return time.Minute
}
