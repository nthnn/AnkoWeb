package std

import (
	"bytes"
	"encoding/json"
)

var lastErrorString string = ""

func MapToJsonFn(data map[string]interface{}) string {
	converted, err := json.Marshal(data)
	if err != nil {
		lastErrorString = err.Error()
		return "{}"
	}

	return string(converted)
}

func JsonToMapFn(jsonStr string) map[string]interface{} {
	result := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &result)

	return result
}

func LastJsonErrorFn() string {
	return lastErrorString
}

func JsonPrettifyFn(jsonStr string, indent string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(jsonStr), "", indent)
	if err != nil {
		lastErrorString = err.Error()
		return jsonStr
	}

	return out.String()
}
