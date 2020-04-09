package Net

import (
	"fmt"
	"main.go/tuuz/Array"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"net/http"
	"net/url"
	"strings"
)

/*
headers = map[string]string{
"User-Agent":    "Sublime",
"Authorization": "Bearer access_token",
"Content-Type":  "application/json",
}

cookies = map[string]string{
"userId":    "12",
"loginTime": "15045682199",
}

queries = map[string]string{
"page": "2",
"act":  "update",
}

postData = map[string]interface{}{
"name":      "mike",
"age":       24,
"interests": []string{"basketball", "reading", "coding"},
"isAdmin":   true,
}
*/

func Post(url string, values map[string]interface{}, headers map[string]string, cookies map[string]string) (int, interface{}, error) {
	// 链式操作
	req := Request()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	ret, err := req.Post(url, values)
	body, err := ret.Body()
	if err != nil {
		return 500, "", err
	} else {
		return 0, string(body), err
	}
}

func PostCookie(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) (int, interface{}, map[string]interface{}, error) {
	req := Request()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	ret, err := req.Post(url+"?"+Http_build_query(queries), postData)
	body, err := ret.Body()
	cookie_arr := CookieHandler(ret.Cookies())
	//fmt.Println(cookie_arr)
	if err != nil {
		return 500, "", cookie_arr, err
	} else {
		return 0, string(body), cookie_arr, err
	}
}

func PostCookieAuto(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, ident string) (float64, interface{}, error) {
	req := Request()
	cookies, err := CookieSelector(ident)
	cook := Array.Mapinterface2MapString(cookies)

	req.SetHeaders(headers)
	req.SetCookies(cook)
	ret, err := req.Post(url+"?"+Http_build_query(queries), postData)
	body, err := ret.Body()

	cookie_arr := CookieHandler(ret.Cookies())
	CookieUpdater(cookie_arr, ident)
	if err != nil {
		return 500, "", err
	} else {
		return 0, string(body), err
	}
}

func PostCookieManual(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookie map[string]interface{}, ident string) (float64, interface{}, error) {
	req := Request()
	CookieUpdater(cookie, ident)
	cookies, err := CookieSelector(ident)
	cook := Array.Mapinterface2MapString(cookies)

	req.SetHeaders(headers)
	req.SetCookies(cook)
	ret, err := req.Post(url+"?"+Http_build_query(queries), postData)
	body, err := ret.Body()

	cookie_arr := CookieHandler(ret.Cookies())
	CookieUpdater(cookie_arr, ident)
	if err != nil {
		return 500, "", err
	} else {
		return 0, string(body), err
	}
}

/*
headers = map[string]string{
"User-Agent":    "Sublime",
"Authorization": "Bearer access_token",
"Content-Type":  "application/json",
}

cookies = map[string]string{
"userId":    "12",
"loginTime": "15045682199",
}

queries = map[string]string{
"page": "2",
"act":  "update",
}
*/
func Get(url string, queries map[string]interface{}, headers map[string]string, cookies map[string]string) (float64, interface{}, error) {
	req := Request()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	ret, err := req.Get(url, queries)
	if err != nil {
		fmt.Println(err)
		return 500, "", err
	}
	body, err := ret.Body()

	if err != nil {
		return 500, "", err
	} else {
		return 0, string(body), err
	}
}

func GetCookie(url string, queries map[string]interface{}, headers map[string]string, cookies map[string]string) (float64, interface{}, map[string]interface{}, error) {
	req := Request()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	ret, err := req.Get(url, queries)
	body, err := ret.Body()
	cookie_arr := CookieHandler(ret.Cookies())
	//fmt.Println(cookie_arr)
	if err != nil {
		return 500, "", cookie_arr, err
	} else {
		return 0, string(body), cookie_arr, err
	}
}

func GetCookieAuto(url string, queries map[string]interface{}, headers map[string]string, ident string) (float64, interface{}, error) {
	req := Request()
	cookies, err := CookieSelector(ident)
	cook := Array.Mapinterface2MapString(cookies)

	req.SetHeaders(headers)
	req.SetCookies(cook)
	ret, err := req.Get(url, queries)
	if err != nil {
		fmt.Println(err)
		return 500, "", err
	}
	body, err := ret.Body()
	cookie_arr := CookieHandler(ret.Cookies())
	CookieUpdater(cookie_arr, ident)
	if err != nil {
		return 500, "", err
	} else {
		return 0, string(body), err
	}
}

func GetCookieManual(url string, queries map[string]interface{}, headers map[string]string, cookie map[string]interface{}, ident string) (float64, interface{}, error) {
	req := Request()
	CookieUpdater(cookie, ident)
	cookies, err := CookieSelector(ident)
	cook := Array.Mapinterface2MapString(cookies)

	req.SetHeaders(headers)
	req.SetCookies(cook)
	ret, err := req.Get(url, queries)
	if err != nil {
		fmt.Println(err)
		return 500, "", err
	}
	body, err := ret.Body()
	cookie_arr := CookieHandler(ret.Cookies())
	CookieUpdater(cookie_arr, ident)
	if err != nil {
		return 500, "", err
	} else {
		return 0, string(body), err
	}
}

func CookieHandler(resp_headers []*http.Cookie) map[string]interface{} {
	cookie_arr := make(map[string]interface{})
	for _, resp_header := range resp_headers {
		cookie_arr[resp_header.Name] = resp_header.Value
	}
	return cookie_arr
}

func CookieHandler2(resp_header map[string]interface{}) map[string]interface{} {
	cookie := strings.Split(Calc.Any2String(resp_header["Set-Cookie"]), "; ")
	cookie_arr := make(map[string]interface{})
	for _, v := range cookie {
		split := strings.Split(v, "=")
		if CookieTagChecker(split[0]) == true {
			cookie_arr[split[0]] = split[1]
		}
	}
	return cookie_arr
}

func CookieUpdater(new_cookie map[string]interface{}, ident string) {
	user_cookie_map, err := CookieSelector(ident)
	if err != nil {
		fmt.Println(err)
		Log.Err(err)
		user_cookie_map = new_cookie
	} else {
		user_cookie_map = Array.Merge(user_cookie_map, new_cookie)
	}
	//_, err = Redis.Set("__cookie__"+ident, user_cookie_map, 30*86400)
	if err != nil {
		fmt.Println(err)
		Log.Err(err)
		return
	}
}

func CookieSelector(ident string) (map[string]interface{}, error) {
	user_cookie_map, err := Redis.Get("__cookie__" + ident)
	if err != nil {
		return make(map[string]interface{}), err
	}
	//fmt.Println(user_cookie_map)
	return user_cookie_map.(map[string]interface{}), err
}

func Http_build_query(querymap map[string]interface{}) string {
	query := make(url.Values)
	for k, v := range querymap {
		query.Add(k, Calc.Any2String(v))
	}
	//fmt.Println(query.Encode())
	return query.Encode()
}
