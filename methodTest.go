package main

import (
	"GoWithTestAutomation/utils"
	"fmt"
)

func main() {
	j, _ := utils.ReadJsonFile("/json/sample.json")
	d, _ := utils.FlatJson2Xml(j, ".")
	s, _ := d.WriteToString()
	fmt.Println(s)
}
