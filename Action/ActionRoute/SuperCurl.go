package ActionRoute

import (
	"fmt"
	"main.go/Conf"
	"main.go/tuuz/Array"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Net"
	"main.go/tuuz/RET"
	"net"
	"time"
)

func Curl(url string, method string, values map[string]interface{}, headers map[string]interface{}, cookie map[string]interface{}, typ string, echo string, conn net.TCPConn, route string, delay float64) {
	ts := Calc.Any2String(delay)
	del, err := time.ParseDuration(ts + "s")
	if err != nil {

	} else {
		time.Sleep(del)
	}
	SuperCurl(url, method, values, headers, cookie, typ, echo, conn, route)
}

func SuperCurl(url string, method string, values map[string]interface{}, headers map[string]interface{}, cookie map[string]interface{}, typ string, echo string, conn net.TCPConn, route string) {
	header := Array.Mapinterface2MapString(headers)
	cookies := Array.Mapinterface2MapString(cookie)
	if method == "post" {
		req := Net.Request()
		req.SetHeaders(header)
		req.SetCookies(cookies)
		ret, err := req.Post(url, values)
		if err != nil {
			ecam("[BiliHP-NET-ERROR0]:", err, "")
			return
		}
		body, err := ret.Body()
		if err != nil {
			ecam("[BiliHP-NET-ERROR1]:", err, "")
			return
		}
		resp_header := ret.Headers()
		//fmt.Println(cookie_arr)
		if err != nil {
			ecam("[BiliHP-NET-ERROR2]:", err, "")
			return
		} else {
			ret := make(map[string]interface{})
			ret["cookie"] = Net.CookieHandler(resp_header)
			ret["route"] = route
			ret["body"] = Jsong.Decode(string(body))
			ret["header"] = resp_header
			ret["statusCode"] = 200
			if Conf.LoadConf("debug", "debug") == "true" {
				fmt.Println("ret-post:", RET.Ws_succ(typ, 0, ret, echo))
			}
			Send(conn, SendObj(typ, ret, echo, values))
		}
	} else {
		req := Net.Request()
		req.SetHeaders(header)
		req.SetCookies(cookies)
		ret, err := req.Get(url, values)
		if err != nil {
			ecam("[BiliHP-NET-ERROR10]:", err, "")
			return
		}
		body, err := ret.Body()
		if err != nil {
			ecam("[BiliHP-NET-ERROR11]:", err, "")
			return
		}
		resp_header := ret.Headers()
		if err != nil {
			ecam("[BiliHP-NET-ERROR12]:", err, "")
			return
		} else {
			ret := make(map[string]interface{})
			ret["cookie"] = Net.CookieHandler(resp_header)
			ret["route"] = route
			ret["body"] = Jsong.Decode(string(body))
			ret["header"] = resp_header
			ret["statusCode"] = 200
			if Conf.LoadConf("debug", "debug") == "true" {
				fmt.Println("ret-get:", RET.Ws_succ(typ, 0, ret, echo))
			}
			Send(conn, RET.Ws_succ(typ, 0, ret, echo))
		}
	}
}
