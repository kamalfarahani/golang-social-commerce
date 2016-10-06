package controllers

import "encoding/json"

func jsonStr(data interface{}) string {
	jData, _ := json.Marshal(data)
	return string(jData)
}
