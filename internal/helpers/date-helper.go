package helpers

import (
	"time"
)

func ParseDate(dateString string) (date time.Time, err error) {
	layout := "2006-01-02"
	return time.Parse(layout, dateString)
}
