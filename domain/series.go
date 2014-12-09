package domain

import (
	"time"
)

type TimeSeries struct {
	When  time.Time
	Posts []*Post
}
