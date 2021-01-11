package utils

import (
	"bytes"
	"fmt"
)

// MapToString converts Map to String for output
func MapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s:\"%s\"\n", key, value)
	}
	return b.String()
}
