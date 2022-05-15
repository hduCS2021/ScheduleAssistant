package date

import "time"

type Weekday uint8

func NewWeekday(weekday string) (Weekday, error) {
	var w Weekday
	err := parseString(&w, 1, 7, weekday)
	return w, err
}

func (Weekday) GetTimeUnit() time.Duration {
	return time.Hour * 24
}

func (w Weekday) IsValid(t time.Time) bool {
	return (w>>((t.Weekday()+6)%7))%2 == 1
}
