package date

import (
	"fmt"
	"strconv"
	"strings"
)

type Minute struct {
	minute uint64
}

func New(minute string) (*Minute, error) {
	res := strings.FieldsFunc(minute, func(r rune) bool {
		return r == ','
	})
	m := Minute{}
	for i := 59; i >= 0; i-- {
		for _, v := range res {
			if strings.ContainsRune(v, '-') {
				//时间区间
				var bg, ed int
				if _, err := fmt.Sscanf(v, "%d-%d", &bg, &ed); err != nil {
					return nil, err
				}
				if i <= ed && i >= bg {
					m.minute++
					break
				}
			} else {
				//时刻
				t, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				if t == i {
					m.minute++
					break
				}
			}
		}
		if i > 0 {
			m.minute <<= 1
		}
	}
	return &m, nil
}

func (m *Minute) IsValid(minute int) bool {
	return (m.minute>>minute)%2 == 1
}

//GetNextMinute return next valid Minute after now(now not included)
//input should be in range -1 to 59
//it returns the first valid Minute, if now is -1
//if there is no valid Minute, -1 will be returned
func (m *Minute) GetNextMinute(now int) int {
	var tmp uint64
	tmp = m.minute >> (now + 1)
	pos := now + 1
	for ; pos < 60; pos++ {
		if tmp%2 == 1 {
			return pos
		}
		tmp >>= 1
	}
	return -1
}
