package date

import "time"

var BeginDate = time.Date(2022, 2, 21, 0, 0, 0, 0, time.Local)

func SetBeginDate(t time.Time) {
	BeginDate = t.Round(time.Hour * 24).Add((time.Duration(t.Weekday()+6) % 7) * (-time.Hour * 24))
}
