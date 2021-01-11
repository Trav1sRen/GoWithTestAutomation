package utils

import (
	"bufio"
	"encoding/json"
	"github.com/beevik/etree"
	"github.com/bitly/go-simplejson"
	"io"
	"os"
	"regexp"
	"strings"
)

// ReadJSONFile returns JSONObject from file at specified path
func ReadJSONFile(path string) (j *simplejson.Json, err error) {
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

// UnflattenJSON converts flat JSONObject to nested JSONObject
func UnflattenJSON(j *simplejson.Json, delim string) (newj *simplejson.Json, err error) {
	newj = simplejson.New()

	var m map[string]interface{}
	if m, err = j.Map(); err != nil {
		return
	}

	for k, v := range m {
		path := strings.Split(k, delim)
		newj.SetPath(path, v)
	}
	return
}

// FlatJSON2XML converts flat JSONObject to ElementTree Document
func FlatJSON2XML(j *simplejson.Json, delim, dupSymbol string) (doc *etree.Document, err error) {
	doc = etree.NewDocument()

	var m map[string]interface{}
	if m, err = j.Map(); err != nil {
		return
	}

	var cur *etree.Element
	cur = &doc.Element
	for k, v := range m {
		path := strings.Split(k, delim)
		for i, tag := range path {
			attrs := make(map[string]string)

			if reg := regexp.MustCompile(`(\w+)\[(.+)]`); reg != nil {
				m := reg.FindAllStringSubmatch(tag, 1)
				if len(m) != 0 {
					tag = m[0][1]
					pairs := m[0][2]
					if reg = regexp.MustCompile(`\w+=\w+`); reg != nil {
						for _, pair := range reg.FindAllString(pairs, -1) {
							p := strings.Split(pair, "=")
							attrs[p[0]] = p[1]
						}
					}
				}
			}

			var f *etree.Element
			createElement := func() {
				cur = cur.CreateElement(tag)
				for k, v := range attrs {
					cur.CreateAttr(k, v)
				}
				attrs = make(map[string]string)
			}

			if tag[len(tag)-1:] == dupSymbol {
				tag = tag[:len(tag)-1]
				f = doc.FindElement("./" + strings.Join(path[:i], "/"))
				cur = f
				createElement()
			} else {
				f = doc.FindElement("./" + strings.Join(path[:i+1], "/"))
				if f != nil {
					cur = f
					continue
				}
				createElement()
			}
		}
		if v.(string) != "" {
			cur.CreateText(v.(string))
		}
	}

	doc.Indent(2)
	return
}

// JSON2Str converts JSONObject to String
func JSON2Str(j *simplejson.Json) (s string, err error) {
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
