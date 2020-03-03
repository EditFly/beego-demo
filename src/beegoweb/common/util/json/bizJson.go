package json

import (
	"encoding/json"
)

func Stringify(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", nil
	}
	//logs.Info(string(b))
	return string(b), nil
}

// json转map int会变为float64 注意
func Parse(str string, obj interface{}) error {
	err := json.Unmarshal([]byte(str), &obj)
	if err != nil {
		return err
	}
	return nil
}
