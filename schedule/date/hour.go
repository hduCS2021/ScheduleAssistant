package date

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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
//input should be in range [0,23]
func (m Hour) IsValid(t time.Time) bool {
	return (m>>t.Hour())%2 == 1
}
