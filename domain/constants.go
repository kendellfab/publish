package domain

import (
	"encoding/json"
	"time"
)

// Mon Jan 2 15:04:05 -0700 MST 2006
const (
	DATE_STORAGE_FORMAT = "2006-01-02T15:04:05Z"
	SLUG_SPACER         = "-"
	ALREADY_EXISTS      = "already exists"
	SESS_FLASH          = "flash"
	FLASH_ERROR         = "error_flash"
	FLASH_SUCCESS       = "success_flash"
)

func SerializeDate(input time.Time) string {
	return input.Format(DATE_STORAGE_FORMAT)
}

func DeserializeDate(input string) (time.Time, error) {
	return time.Parse(DATE_STORAGE_FORMAT, input)
}

func SerializeTags(input []string) (string, error) {
	val, err := json.Marshal(input)
	return string(val), err
}

func DeserializeTags(input string) ([]string, error) {
	var tags []string
	err := json.Unmarshal([]byte(input), tags)
	return tags, err
}

func DateComponents(input time.Time) (int, int, int) {
	return input.Day(), int(input.Month()), input.Year()
}
