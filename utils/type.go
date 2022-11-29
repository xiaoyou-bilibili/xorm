package utils

import "encoding/json"

func Interface2String(data interface{}) string {
	if res, err := json.Marshal(data); err != nil {
		return ""
	} else {
		return string(res)
	}
}
