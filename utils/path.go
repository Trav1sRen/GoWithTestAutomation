package utils

import (
	"os"
)

// getRootPath returns the root path of current project
func getRootPath() (s string) {
	s, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return
}
