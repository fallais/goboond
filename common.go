package goboond

import (
	"fmt"
	"strings"
	"time"
)

const DateFormat = "2006-01-02"

type SearchOptions struct {
	MaxResults int    `json:"maxResults"`
	Page       int    `json:"page"`
	Sort       string `json:"sort"`
}

func formatIntArray(arr []int) string {
	if len(arr) == 0 {
		return ""
	}
	// Use fmt.Sprint to get the array as a string
	return fmt.Sprintf("[%s]", strings.Trim(strings.Replace(fmt.Sprint(arr), " ", ",", -1), "[]"))
}

func formatStringArray(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	// Use fmt.Sprint to get the array as a string
	return fmt.Sprintf("[%s]", strings.Trim(strings.Replace(fmt.Sprint(arr), " ", ",", -1), "[]"))
}

type Period string

const (
	PeriodCreated Period = "created"
	PeriodStarted Period = "started"
	PeriodUpdated Period = "updated"
	PeriodRunning Period = "running"
)

// StringPtr returns a pointer to the string passed as argument.
func StringPtr(s string) *string {
	return &s
}

// BoolPtr returns a pointer to the bool passed as argument.
func BoolPtr(b bool) *bool {
	return &b
}

// TimePtr returns a pointer to the time passed as argument.
func TimePtr(b time.Time) *time.Time {
	return &b
}

func PeriodPtr(b Period) *Period {
	return &b
}
