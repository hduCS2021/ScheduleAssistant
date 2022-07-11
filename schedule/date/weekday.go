package date

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Weekday uint8

func NewWeekday(weekday string) Weekday {
	var w Weekday
	if err := parseString(&w, 1, 7, weekday); err != nil {
		log.Warnf("fail to parse time:%v", err)
	}
	return w
}

func (Weekday) GetTimeUnit() time.Duration {
	return time.Hour * 24
}

func (w Weekday) IsValid(t time.Time) bool {
	return (w>>((t.Weekday()+6)%7))%2 == 1
}
