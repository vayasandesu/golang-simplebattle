package jsonexample

import (
	"encoding/json"
)

func Serialize(obj JsonStruct) string {
	text, err := json.Marshal(&obj)
	if err != nil {
		return ""
	}

	return string(text)
}

func Deserialize(text string) JsonStruct {
	var value JsonStruct
	err := json.Unmarshal([]byte(text), &value) //use &value because we want to write result to value variable
	if err != nil {
		return JsonStruct{}
	}

	return value
}
