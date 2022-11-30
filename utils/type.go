package utils

import (
	"encoding/json"
	"log"
)

func Interface2String(data interface{}) string {
	if res, err := json.Marshal(data); err != nil {
		log.Printf("marshal data err %v", err)
		return ""
	} else {
		return string(res)
	}
}
