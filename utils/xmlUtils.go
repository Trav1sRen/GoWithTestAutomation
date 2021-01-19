package utils

import (
	"github.com/clbanning/mxj/v2"
)

// XML2Map converts xml string to an iterable map as type map[string]interface{}
// mxj will auto-trim the namespace of tags
// tag attributes start with "-" as the key in map
func XML2Map(xmlStr string) (m map[string]interface{}, err error) {
	return mxj.NewMapXml([]byte(xmlStr))
}
