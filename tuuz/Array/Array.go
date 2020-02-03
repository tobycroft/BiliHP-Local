package Array

import (
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
)

func Merge(args ...map[string]interface{}) map[string]interface{} {
	arr := make(map[string]interface{})
	for _, arrs := range args {
		for key, value := range arrs {
			arr[key] = value
		}
	}
	return arr
}

func Mapinterface2MapString(maps map[string]interface{}) map[string]string {
	strs := make(map[string]string)
	for key, value := range maps {
		switch value.(type) {
		case string:
			strs[key] = value.(string)
			break
		case int:
			tmp := value.(int)
			strs[key] = Calc.Int2String(tmp)
			break
		case int64:
			tmp := value.(int64)
			strs[key] = Calc.Int642String(tmp)
			break
		case float64:
			tmp := value.(float64)
			strs[key] = Calc.Float642String(tmp)
			break

		case float32:
			tmp := value.(float64)
			strs[key] = Calc.Float642String(tmp)
			break

		default:
			strs[key] = value.(string)
			break
		}

	}
	return strs
}

func MapString2MapInterface(maps map[string]string) map[string]interface{} {
	arr := make(map[string]interface{})
	for key, value := range maps {
		arr[key] = value
	}
	return arr
}

func MapString2Interface(maps map[string]interface{}) map[string]interface{} {
	strs := make(map[string]interface{})
	for key, value := range maps {
		switch value.(type) {
		case string:
			strs[key] = value.(string)
		case int:
			tmp := value.(int)
			strs[key] = Calc.Int2String(tmp)
		case int64:
			tmp := value.(int64)
			strs[key] = Calc.Int642String(tmp)
		}
	}
	return strs
}
func InArray(str interface{}, haystack []interface{}) bool {
	str, _ = Jsong.Encode(str)
	for _, v := range haystack {
		v, _ = Jsong.Encode(v)
		if str == v {
			return true
		}
	}
	return false
}
