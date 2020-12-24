package utils

import (
	"os"
)

func getRootPath() (s string) {
	s, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return
}
