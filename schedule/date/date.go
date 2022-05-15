package date

import (
	"sort"
	"time"
)

type Date interface {
	// GetTimeUnit returns the Length of a type of time.
	GetTimeUnit() time.Duration
	IsValid(t time.Time) bool
}

type List struct {
	dates []Date
}

func (l *List) Append(d Date) {
	l.dates = append(l.dates, d)
}

func (l *List) sort() {
	sort.Slice(l.dates, func(i, j int) bool {
		return l.dates[i].GetTimeUnit() > l.dates[i].GetTimeUnit()
	})
}

func (l *List) Check(t time.Time) bool {
	for _, v := range l.dates {
		if !v.IsValid(t) {
			return false
		}
	}
	return true
}
