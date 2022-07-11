package date

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type WeekPuppet uint8

// NewWeekPuppet
// 0:none week
// 1:single week
// 2:double week
// 3:both week
func NewWeekPuppet(puppet int) WeekPuppet {
	if puppet >= 0 && puppet <= 3 {
		return WeekPuppet(puppet)
	}
	log.Warnf("fail to parse time:invalid weekPuppet %d", puppet)
	return 0
}

func (w WeekPuppet) IsValid(t time.Time) bool {
	dur := t.Sub(BeginDate)
	dur /= time.Hour * 24 * 7
	return (w>>(dur%2))%2 == 1
}

func (w WeekPuppet) GetTimeUnit() time.Duration {
	return time.Hour * 24 * 7
}
