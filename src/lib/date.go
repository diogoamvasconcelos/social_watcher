package lib

import (
	"fmt"
	"time"
)

func ToISO8061(date time.Time) string {
	return fmt.Sprintf(date.Format(time.RFC3339))
}

func MinutesAgo(minutes int) time.Time {
	return time.Now().Add(time.Duration(-minutes) * time.Minute)
}
