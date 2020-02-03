package RET

import (
	"fmt"
	"main.go/tuuz/Jsong"
)

func Json(data interface{}) string {
	ret, _ := Jsong.Encode(data)
	return ret
}

func Success(code interface{}, data interface{}) string {
	ret := make(map[string]interface{})
	ret["code"] = code
	ret["data"] = data
	jb, err := Jsong.Encode(ret)
	//fmt.Println(jb)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jb)
}

func Fail(code int, data interface{}) interface{} {
	return Success(code, data)
}

func Ret_succ(code interface{}, data interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	ret["code"] = code
	ret["data"] = data
	return ret
}

func Ret_fail(code interface{}, data interface{}) map[string]interface{} {
	return Ret_succ(code, data)
}

func Ws_succ(typ string, code interface{}, data interface{}, echo interface{}) string {
	ret := make(map[string]interface{})
	ret["type"] = typ
	ret["code"] = code
	ret["data"] = data
	ret["echo"] = echo
	jb, err := Jsong.Encode(ret)
	//fmt.Println(jb)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jb)
}

func Ws_fail(typ string, code interface{}, data interface{}, echo interface{}) string {
	return Ws_succ(typ, code, data, echo)
}
