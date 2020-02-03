package Net

import "strings"

const cookie_tag = "sid,JSESSIONID,DedeUserID,DedeUserID__ckMd5,SESSDATA,bili_jct,sid"

func CookieTagChecker(cookie_key string) bool {
	if cookie_tag == "" {
		return true
	} else {
		arr := strings.Split(cookie_tag, ",")
		for _, v := range arr {
			if v == cookie_key {
				return true
			}
		}
		return false
	}
}
