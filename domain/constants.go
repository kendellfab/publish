package domain

import (
	"encoding/json"
	"time"
)

const (
	ACTIVE         = "active"
	SLUG_SPACER    = "-"
	ALREADY_EXISTS = "already exists"
	SESS_AUTH_KEY  = "sess_auth_key"
	SESS_USER_ID   = "sess_user_id"
	SESS_LOGGED_IN = "sess_logged_in"
	SESS_FLASH     = "flash"
	FLASH_ERROR    = "error_flash"
	FLASH_SUCCESS  = "success_flash"
	CONTEXT_USER   = "cntxt_user"
)

func SerializeDate(input time.Time) string {
	return input.Format(time.RFC3339)
}

func DeserializeDate(input string) (time.Time, error) {
	return time.Parse(time.RFC3339, input)
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
