package Jsong

func ParseObject(data interface{}) (map[string]interface{}, error) {
	ret, err := Encode(data)
	if err != nil {
		return nil, err
	}
	return JObject(ret)
}

func ParseSlice(data interface{}) ([]interface{}, error) {
	ret, err := Encode(data)
	if err != nil {
		return nil, err
	}
	return JArray(ret)
}
