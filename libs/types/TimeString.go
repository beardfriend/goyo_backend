package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type TimeString time.Time

func (t TimeString) Value() (driver.Value, error) {
	h := time.Time(t).Hour()

	min := time.Time(t).Minute()
	s := time.Time(t).Second()
	y, m, d := time.Time(t).Date()
	return time.Date(y, m, d, h, min, s, 0, time.Time(t).Location()), nil
}

func (t *TimeString) Scan(value any) error {
	switch value.(type) {
	case time.Time:
		*t = TimeString(value.(time.Time))
		break
	default:
		parsed, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s", value))
		*t = TimeString(parsed)
		break
	}
	return nil
}

func (t TimeString) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02T15:04:05"))
	return []byte(stamp), nil
}
