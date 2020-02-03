package ActionRoute

import (
	"fmt"
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
		body, err := ret.Body()
		resp_header := ret.Headers()
		//fmt.Println(cookie_arr)
		if err != nil {
			ecam("[BiliHP-NET-ERROR]:", err, "")
			return
		} else {
			ret := make(map[string]interface{})
			ret["cookie"] = Net.CookieHandler(resp_header)
			ret["route"] = route
			ret["body"] = Jsong.Decode(string(body))
			ret["header"] = resp_header
			ret["statusCode"] = 200
			fmt.Println("ret:", RET.Ws_succ(typ, 0, ret, echo))
			Send(conn, RET.Ws_succ(typ, 0, ret, echo))
		}
	} else {
		req := Net.Request()
		req.SetHeaders(header)
		req.SetCookies(cookies)
		ret, err := req.Get(url, values)
		body, err := ret.Body()
		resp_header := ret.Headers()
		if err != nil {
			ecam("[BiliHP-NET-ERROR]:", err, "")
			return
		} else {
			ret := make(map[string]interface{})
			ret["cookie"] = Net.CookieHandler(resp_header)
			ret["route"] = route
			ret["body"] = Jsong.Decode(string(body))
			ret["header"] = resp_header
			ret["statusCode"] = 200
			fmt.Println("ret:", RET.Ws_succ(typ, 0, ret, echo))
			Send(conn, RET.Ws_succ(typ, 0, ret, echo))
		}
	}
}
