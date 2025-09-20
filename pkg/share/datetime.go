package share

import (
	"time"
)

func ParseDate(dateStr string) *time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic(err)
	}
	return &t
}
