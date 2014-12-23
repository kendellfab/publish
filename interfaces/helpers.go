package interfaces

import (
	"strings"
)

func GetPlaceholders(items []interface{}) string {
	size := len(items)
	ph := make([]string, size)
	for i := 0; i < size; i++ {
		ph[i] = "?"
	}
	return strings.Join(ph, ",")
}
