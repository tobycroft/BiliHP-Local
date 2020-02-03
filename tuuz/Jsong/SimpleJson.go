package Jsong

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func Simple(json string) (*simplejson.Json, error) {
	res, err := simplejson.NewJson([]byte(json))
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	return res, err
}

func SimpleDecode(json string) *simplejson.Json {
	ret, _ := Simple(json)
	return ret
}
