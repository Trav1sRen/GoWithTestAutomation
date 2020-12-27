package utils

import (
	"bufio"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"io"
	"os"
	"strings"
)

func ReadJsonFile(path string) (j *simplejson.Json, err error) {
	defPath := getRootPath() + path

	var input *os.File
	input, err = os.Open(defPath)
	if err != nil {
		return
	}
	defer input.Close()

	var s string
	reader := bufio.NewReader(input)
	for {
		line, err := reader.ReadString('\n')
		s += line
		if err == io.EOF {
			break
		}
	}

	j, err = simplejson.NewJson([]byte(s))
	if err != nil {
		return
	}
	return
}

func UnflattenJson(j *simplejson.Json) (newj *simplejson.Json, err error) {
	newj = simplejson.New()

	var m map[string]interface{}
	if m, err = j.Map(); err != nil {
		return
	}

	for k, v := range m {
		path := strings.Split(k, ".")
		newj.SetPath(path, v)
	}
	return
}

func Json2Str(j *simplejson.Json) (s string, err error) {
	var m map[string]interface{}
	if m, err = j.Map(); err != nil {
		return
	}

	var b []byte
	if b, err = json.Marshal(m); err != nil {
		return
	}
	s = string(b)
	return
}
