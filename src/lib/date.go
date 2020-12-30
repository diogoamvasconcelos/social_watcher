package lib

import (
	"fmt"
	"time"
)

func ToISO8061(date time.Time) string {
	return fmt.Sprintf(date.Format(time.RFC3339))
}

func FromISO8061(str string) (time.Time, error) {
	date, err := time.Parse(time.RFC3339, str)
	return date, err
}

func MinutesAgo(minutes int) time.Time {
	return time.Now().Add(time.Duration(-minutes) * time.Minute)
}
