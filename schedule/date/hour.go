package date

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Hour uint32

func NewHour(hour string) Hour {
	var h Hour
	if err := parseString(&h, 0, 23, hour); err != nil {
		log.Warnf("fail to parse time:%v", err)
	}
	return h
}

//IsValid check if Hour in range
//input should be in range [0,23]
func (m Hour) IsValid(t time.Time) bool {
	return (m>>t.Hour())%2 == 1
}

func (m Hour) GetTimeUnit() time.Duration {
	return time.Hour
}
