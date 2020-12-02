package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

func Must(e error) {
	if e != nil {
		panic(e)
	}
}

func MustNil(e error) {

}
func StructToJson(object interface{}) (string, error) {
	str, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func StructToJsonOrError(object interface{}) (string) {
	str, err := json.Marshal(object)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

func PrintStrcut(obj interface{}) {
	str, _ := json.Marshal(obj)
	fmt.Println(string(str))
}
func NowSecond() int32 {
	return int32(time.Now().Unix())
}
func NowMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func StructToMap(s interface{}) (map[string]interface{}, error) {
	buf, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	r := map[string]interface{}{}
	err = json.NewDecoder(bytes.NewReader(buf)).Decode(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
