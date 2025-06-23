package parser

import (
	"time"
)

func ParseDateTime(x *time.Time) *string {
	if x == nil {
		return nil
	}
	formattedDate := x.String()
	if formattedDate == "" {
		return nil
	}
	return &formattedDate
}

func ParseTimeTime(x *string) *time.Time {
	if x == nil {
		return nil
	}
	parsedTime, err := time.Parse(time.RFC3339, *x)
	if err != nil {
		return nil // or handle the error as needed
	}
	return &parsedTime
}
