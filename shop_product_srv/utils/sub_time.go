package utils

import (
	"fmt"
	"time"
)

func AddHour(h int) string {
	// h = 24
	now := time.Now()
	after_hour := fmt.Sprintf("+%dh", h)
	duration_time, _ := time.ParseDuration(after_hour)
	return now.Add(duration_time).Format("2006-01-02 15:04:05")
}
