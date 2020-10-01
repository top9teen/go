package utils

import (
	"encoding/json"

	"entitys"
)

func Response_Handle(data map[string]interface{}, er error) (map[string]interface{}, error) {
	var res entitys.Response
	if er == nil {
		res.StatusCode = 200
		res.StatusDesc = "success"
		res.Body = data

	} else {
		res.StatusCode = 302
		res.StatusDesc = "Fail"
		res.Body = data

	}

	var content map[string]interface{}
	json_byte, err := json.Marshal(res)
	json.Unmarshal(json_byte, &content)

	return content, err
}
