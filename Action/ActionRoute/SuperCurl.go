package ActionRoute

import (
	"main.go/Tcp"
	"main.go/tuuz/Array"
	"main.go/tuuz/Net"
	"main.go/tuuz/RET"
	"net"
)

func Curl(url string, method string, values map[string]interface{}, headers map[string]interface{}, cookie map[string]interface{}, typ string, echo string, conn net.TCPConn, route string, delay float64) {

}

func SuperCurl(username string, url string, method string, values map[string]interface{}, headers map[string]interface{}, cookie map[string]interface{}, typ string, echo string, conn net.TCPConn, route string) {
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
			ret[""]
			Tcp.Send(conn, RET.Ws_succ(typ, 0, ret, echo))
		}
	} else {

	}
}
