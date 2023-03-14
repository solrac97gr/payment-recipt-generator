package utils

import (
	"encoding/json"
	"fmt"
)

func PrintIdentJSON(data interface{}) {
	jdata, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(jdata))
}
