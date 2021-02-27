package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// MapToString converts Map to String for output
func MapToString(m map[string]interface{}) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s:\"%s\"\n", key, value)
	}
	return b.String()
}

// MapToJSON converts Map to JSONObject
func MapToJSONStr(m map[string]interface{}) (s string, err error) {
	var b []byte
	if b, err = json.Marshal(m); err != nil {
		return
	}
	s = string(b)
	return
}
