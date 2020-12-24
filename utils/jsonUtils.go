package utils

import (
	"bufio"
	"github.com/bitly/go-simplejson"
	"io"
	"log"
	"os"
)

func readJsonFile(path string) (json *simplejson.Json, err error) {
	defPath := getRootPath() + path

	var input *os.File
	input, err = os.Open(defPath)
	if err != nil {
		log.Printf("Error when opening the file '%s': %v", defPath, err)
		return
	}
	defer input.Close()

	var s string
	reader := bufio.NewReader(input)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		s += line
	}

	json, err = simplejson.NewJson([]byte(s))
	if err != nil {
		log.Print("Error when parsing bytes to Json: ", err)
		return
	}
	return
}

func UnflattenJson(json *simplejson.Json) (newJson *simplejson.Json, err error) {
	newJson = simplejson.New()

	return
}
