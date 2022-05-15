package date

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Minute uint64

func NewMinute(minute string) (Minute, error) {
	res := strings.FieldsFunc(minute, func(r rune) bool {
		return r == ','
	})
	m := Minute(0)
	for i := 59; i >= 0; i-- {
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

//IsValid check if a minute in range
//input should be in range 0-59
func (m Minute) IsValid(t time.Time) bool {
	return (m>>t.Minute())%2 == 1
}

func (m Minute) GetTimeUnit() time.Duration {
	return time.Minute
}
