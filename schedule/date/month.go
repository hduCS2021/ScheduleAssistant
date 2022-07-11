package date

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Month uint16

func NewMonth(month string) Month {
	var m Month
	if err := parseString(&m, 1, 12, month); err != nil {
		log.Warnf("fail to parse time:%v", err)
	}
	return m
}

func (d Month) IsValid(t time.Time) bool {
	return (d>>(t.Month()-1))%2 == 1
}

func (d Month) GetTimeUnit() time.Duration {
	return time.Hour * 24 * 30
}
