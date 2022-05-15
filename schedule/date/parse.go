package date

import (
	"fmt"
	"strconv"
	"strings"
)

type bitStorage interface {
	Minute | Hour | Day | Month | Weekday
}

func parseString[K bitStorage](m *K, bg, ed int, str string) error {
	res := strings.FieldsFunc(str, func(r rune) bool {
		return r == ','
	})
	for i := ed; i >= bg; i-- {
		for _, v := range res {
			if strings.ContainsRune(v, '-') {
				var bg, ed int
				if _, err := fmt.Sscanf(v, "%d-%d", &bg, &ed); err != nil {
					return err
				}
				if i <= ed && i >= bg {
					*m++
					break
				}
			} else {
				t, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				if t == i {
					*m++
					break
				}
			}
		}
		if i > bg {
			*m <<= 1
		}
	}
	return nil
}
