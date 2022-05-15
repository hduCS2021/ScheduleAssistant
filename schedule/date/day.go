package date

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Day uint32

func NewDay(day string) (Day, error) {
	res := strings.FieldsFunc(day, func(r rune) bool {
		return r == ','
	})
	m := Day(0)
	for i := 31; i > 0; i-- {
		for _, v := range res {
			if strings.ContainsRune(v, '-') {
				//日期区间
				var bg, ed int
				if _, err := fmt.Sscanf(v, "%d-%d", &bg, &ed); err != nil {
					return 0, err
				}
				if i <= ed && i >= bg {
					m++
					break
				}
			} else {
				//日期
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
		if i > 1 {
			m <<= 1
		}
	}
	return m, nil
}

func (d Day) IsValid(t time.Time) bool {
	return (d>>(t.Day()-1))%2 == 1
}

func (d Day) GetTimeUnit() time.Duration {
	return time.Hour * 24
}
