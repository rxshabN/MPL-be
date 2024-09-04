package utils

import (
	"fmt"
	"time"
)

// Timer struct to hold the timer's state
type Timer struct {
	startTime time.Time
	duration  time.Duration
}

// Global variable to hold the timer instance
var GlobalTimer *Timer

// CreateTimer initializes the timer and returns a Timer struct
func CreateTimer(duration time.Duration) *Timer {
	GlobalTimer = &Timer{
		startTime: time.Now(),
		duration:  duration,
	}
	return GlobalTimer
}

// TimeLeft returns the remaining time of the timer without restarting it
func (t *Timer) TimeLeft() int {
	remaining := t.duration - time.Since(t.startTime)
	if remaining <= 0 {
		fmt.Println("Timer has expired!")
		return 0
	} else {
		fmt.Printf("Time remaining: %v\n", remaining.Round(time.Second))
		return int(remaining.Round(time.Second))
	}
}
