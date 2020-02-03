package Calc

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

func Chop(s string, character_mask string) string {
	return strings.TrimRight(s, character_mask)
}

func Any2String(any interface{}) string {
	var str string
	switch any.(type) {
	case string:
		str = any.(string)

	case int:
		tmp := any.(int)
		str = Int2String(tmp)

	case int32:
		tmp := int64(any.(int32))
		str = Int642String(tmp)

	case int64:
		tmp := any.(int64)
		str = Int642String(tmp)

	case float64:
		tmp := any.(float64)
		str = Float642String(tmp)

	case float32:
		tmp := float64(any.(float32))
		str = Float642String(tmp)

	case *big.Int:
		tmp := any.(*big.Int)
		str = tmp.String()

	case nil:
		str = ""

	default:
		fmt.Println(reflect.TypeOf(any))
		str = ""
	}
	return str
}

func String2Int(str string) (int, error) {
	return strconv.Atoi(str)
}

func String2Int64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func String2Float64(str string) (float64, error) {
	float, err := strconv.ParseFloat(str, 64)
	return float, err
}
