package Jsong

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strings"
)

func Encode(data interface{}) (string, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jb, err := json.Marshal(data)
	//fmt.Println(jb)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(jb), err
}

func Decode(data string) interface{} {
	ret, _ := JToken(data)
	return ret
}

func JArray(data string) ([]interface{}, error) {
	var arr []interface{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(data), &arr)
	if err != nil {
		return nil, err
	}
	return arr, err
}

func JObject(data string) (map[string]interface{}, error) {
	var arr map[string]interface{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(data), &arr)
	if err != nil {
		return nil, err
	}
	return arr, err
}

func JToken(data string) (interface{}, error) {
	var arr interface{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(data), &arr)
	if err != nil {
		return nil, err
	}
	return arr, err
}

func TCPJObject(data string) ([]map[string]interface{}, error) {
	var arr []map[string]interface{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//var strs []string
	//fmt.Println(data)
	strs := strings.Split(data, "}{")
	if len(strs) > 1 {
		for i, v := range strs {
			if i == 0 {
				v += "}"
			} else if len(strs) == i+1 {
				v = "{" + v
			} else {
				v = "{" + v + "}"
			}
			err := json.Unmarshal([]byte(v), &arr)
			ret, err := JObject(v)
			if err != nil {
				return nil, err
			} else {
				arr = append(arr, ret)
			}
		}
		return arr, nil
	}
	ret, err := JObject(data)
	if err != nil {
		return arr, err
	} else {
		arr = append(arr, ret)
		return arr, err
	}
	return arr, nil
}

func TCPJArray(data string) ([]interface{}, error) {
	var arr []interface{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var strs []string
	strs = strings.Split(data, "}{")
	if len(strs) > 1 {
		for i, v := range strs {
			if i == 0 {
				v += "}"
			} else if len(strs) == i+1 {
				v = "{" + v
			} else {
				v = "{" + v + "}"
			}
			err := json.Unmarshal([]byte(v), &arr)
			ret, err := JArray(v)
			if err != nil {
				return nil, err
			} else {
				arr = append(arr, ret)
			}
		}
		return arr, nil
	}
	ret, err := JArray(data)
	if err != nil {
		return arr, err
	} else {
		arr = append(arr, ret)
		return arr, err
	}
	return arr, err
}
