package usecases

import (
	"time"
)

// Mon Jan 2 15:04:05 MST 2006

func FormatDate(input time.Time) string {
	return input.Format("Mon Jan 2, 2006")
}

func FormatBool(input bool) string {
	if input {
		return "Yes"
	}
	return "No"
}
