package utils

import (
	"fmt"
	"time"
)

func CreateTimer(duration time.Duration) func() int {
	startTime := time.Now()
	endTime := startTime.Add(duration)

	return func() int {
		remaining := endTime.Sub(time.Now())
		if remaining <= 0 {
			fmt.Println("Timer has expired!")
		} else {
			fmt.Printf("Time remaining: %v\n", remaining.Round(time.Second))
		}
		return int(remaining.Round(time.Second))
	}
}
