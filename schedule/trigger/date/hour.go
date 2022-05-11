package date

import (
	"fmt"
	"strconv"
	"strings"
)

type Hour uint32

func NewHour(hour string) (Hour, error) {
	res := strings.FieldsFunc(hour, func(r rune) bool {
		return r == ','
	})
	m := Hour(0)
	for i := 23; i >= 0; i-- {
		for _, v := range res {
			if strings.ContainsRune(v, '-') {
				//时间区间
				var bg, ed int
				if _, err := fmt.Sscanf(v, "%d-%d", &bg, &ed); err != nil {
					return 0, err
				}
				if i <= ed && i >= bg {
					m++
					break
				}
			} else {
				//时刻
				t, err := strconv.Atoi(v)
				if err != nil {
					return 0, err
				}
				if t == i {
					m++
					break
				}
			}
		}
		if i > 0 {
			m <<= 1
		}
	}
	return m, nil
}

//IsValid check if Hour in range
//input should be in range 0-23
func (m Hour) IsValid(hour int) bool {
	return (m>>hour)%2 == 1
}

//GetNextHour return next valid Hour after now(now not included)
//input should be in range -1 to 23
//it returns the first valid Hour, if now is -1
//if there is no valid Hour, -1 will be returned
func (m Hour) GetNextHour(now int) int {
	var tmp uint32
	tmp = uint32(m) >> (now + 1)
	pos := now + 1
	for ; pos < 24; pos++ {
		if tmp%2 == 1 {
			return pos
		}
		tmp >>= 1
	}
	return -1
}
