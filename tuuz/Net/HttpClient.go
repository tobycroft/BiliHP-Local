package Net

import "github.com/kirinlabs/HttpRequest"

//
func Request() *HttpRequest.Request {
	req := HttpRequest.NewRequest()
	//req := HttpRequest.NewRequest().Debug(true).DisableKeepAlives(false).SetTimeout(5)
	return req
}

//
//func head() {
//	req.SetHeaders(map[string]string{
//		"Content-Type": "application/x-www-form-urlencoded",
//		"Connection":   "keep-alive",
//	})
//
//	req.SetHeaders(map[string]string{
//		"Source": "api",
//	})
//}
//
//func cook() {
//	req.SetCookies(map[string]string{
//		"name":  "json",
//		"token": "",
//	})
//
//	req.SetCookies(map[string]string{
//		"age": "19",
//	})
//}
//
//func auth() {
//	req.SetBasicAuth("username", "password")
//}
//
//func timeout() {
//	req.SetTimeout(5)
//}
//
//func ignorehttptlscert() {
//	req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
//}
//
//func Get() {
//	res, err := req.Get("http://127.0.0.1:8000")
//	res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest")
//	res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest", nil)
//	res, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest", "address=beijing")
//
//	res, err := HttpRequest.Get("http://127.0.0.1:8000")
//	res, err := HttpRequest.Debug(true).SetHeaders(map[string]string{}).Get("http://127.0.0.1:8000")
//	body, err := res.Body()
//	if err != nil {
//		return
//	}
//
//	return string(body)
//}
//
//func aaa() {
//	req := HttpRequest.NewRequest().
//		Debug(true).
//		SetHeaders(map[string]string{
//			"Content-Type": "application/x-www-form-urlencoded",
//		}).SetTimeout(5)
//	res, err := HttpRequest.NewRequest().Get("http://127.0.0.1")
//}
//
//func upload() {
//	res, err := req.Upload("http://127.0.0.1:8000/upload", "/root/demo.txt", "uploadFile")
//	body, err := res.Body()
//	if err != nil {
//		return
//	}
//	return string(body)
//}
//
//func post() {
//	req:=Request();
//	res, err := req.Post("http://127.0.0.1:8000")
//	res, err := req.Post("http://127.0.0.1:8000", "title=github&type=1")
//	res, err := req.JSON().Post("http://127.0.0.1:8000", "{\"id\":10,\"title\":\"HttpRequest\"}")
//	res, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
//		"id":    10,
//		"title": "HttpRequest",
//	})
//	body, err := res.Body()
//	if err != nil {
//		return
//	}
//	return string(body)
//
//	res, err := HttpRequest.Post("http://127.0.0.1:8000")
//	res, err := HttpRequest.JSON().Post("http://127.0.0.1:8000", map[string]interface{}{"title": "github"})
//	res, err := HttpRequest.Debug(true).SetHeaders(map[string]string{}).JSON().Post("http://127.0.0.1:8000", "{\"title\":\"github\"}")
//}
