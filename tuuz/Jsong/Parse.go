package Jsong

import jsoniter "github.com/json-iterator/go"

func ParseObject(data interface{}) (map[string]interface{}, error) {
	ret, err := Encode(data)
	if err != nil {
		return nil, err
	}
	return JObject(ret)
}
func ParseObject2(data interface{}) (map[string]string, error) {
	ret, err := Encode(data)
	if err != nil {
		return nil, err
	}
	var arr map[string]string
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err2 := json.Unmarshal([]byte(ret), &arr)
	if err2 != nil {
		return nil, err
	}
	return arr, err
}
func ParseSlice(data interface{}) ([]interface{}, error) {
	ret, err := Encode(data)
	if err != nil {
		return nil, err
	}
	return JArray(ret)
}
