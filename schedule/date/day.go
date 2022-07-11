package date

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Day uint32

func NewDay(day string) Day {
	var d Day
	if err := parseString(&d, 1, 31, day); err != nil {
		log.Warnf("fail to parse time:%v", err)
	}
	return d
}

func (d Day) IsValid(t time.Time) bool {
	return (d>>(t.Day()-1))%2 == 1
}

func (d Day) GetTimeUnit() time.Duration {
	return time.Hour * 24
}
