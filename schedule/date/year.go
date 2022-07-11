package date

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Year []int

func NewYear(year string) Year {
	var m Year
	res := strings.FieldsFunc(year, func(r rune) bool {
		return r == ','
	})
	for _, v := range res {
		if strings.ContainsRune(v, '-') {
			var bg, ed int
			if _, err := fmt.Sscanf(v, "%d-%d", &bg, &ed); err != nil {
				log.Warnf("fail to parse time: %v", err)
				return m
			}
			for i := bg; i <= ed; i++ {
				if m.isNotExist(i) {
					m = append(m, i)
					sort.Ints(m)
				}
			}
		} else {
			t, err := strconv.Atoi(v)
			if err != nil {
				log.Warnf("fail to parse time: %v", err)
				return m
			}
			if m.isNotExist(t) {
				m = append(m, t)
				sort.Ints(m)
			}
		}
	}
	return m
}

func (m *Year) isNotExist(year int) bool {
	bg := 0
	ed := len(*m) - 1
	for bg <= ed {
		if (*m)[(bg+ed)/2] == year {
			return false
		} else if year < (*m)[(bg+ed)/2] {
			ed = (bg+ed)/2 - 1
		} else {
			bg = (bg+ed)/2 + 1
		}
	}
	return true
}

func (m Year) GetTimeUnit() time.Duration {
	return time.Hour * 24 * 365
}

func (m Year) isValid(t time.Time) bool {
	return !m.isNotExist(t.Year())
}
