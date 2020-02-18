package main

import (
	"fmt"
	"html/template"
	"main.go/Conf"
	"main.go/Tcp"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Net"
	"net/http"
	"time"
)

func main() {
	username := Conf.LoadConf("user", "username")
	token := Conf.LoadConf("user", "token")

	fmt.Println(token)
	if username != "" && token != "" {
		go Tcp.Create(username, token)
	} else {

	}
	//exec.Command(`cmd`, `/c`, `start`, `http://localhost/`).Start()
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/panel", panel)
	http.HandleFunc("/writelogin", writelogin)
	http.HandleFunc("/user/login", UserLoginHandler)
	http.HandleFunc("/setting_get", setting_get)
	http.HandleFunc("/setting_set", setting_set)
	// 设置静态目录
	fsh := http.FileServer(http.Dir("./html"))
	http.Handle("/html/", http.StripPrefix("/html/", fsh))
	fmt.Println("正在启动程序，请访问http://127.0.0.1")
	time.Sleep(5 * time.Second)
	if username == "" || token == "" {
		fmt.Println("你还没有登录，请访问上面的地址进行登录")
	}
	err := http.ListenAndServe("0.0.0.0:80", nil)

	if err != nil {
		fmt.Println("80端口被占用，正在使用81端口重试")
		fmt.Println("正在更换端口并启动程序，请访问http://127.0.0.1:81")
		time.Sleep(5 * time.Second)
		err := http.ListenAndServe("0.0.0.0:79", nil)
		if err != nil {
			fmt.Println("79端口也被占用……程序自动停止")
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	username := Conf.LoadConf("user", "username")
	token := Conf.LoadConf("user", "token")
	if username == "" || token == "" {
		url := "/login"
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		url := "/panel"
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func writelogin(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Handler Hello")
	username := request.PostFormValue("username")
	token := request.PostFormValue("token")
	Conf.SaveConf("user", "username", username)
	Conf.SaveConf("user", "token", token)
	go Tcp.Create(username, token)
	url := "/panel"
	http.Redirect(w, request, url, http.StatusFound)
}

func UserLoginHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Handler Hello")
	fmt.Fprintf(w, "Login Success")
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	username := Conf.LoadConf("user", "username")
	token := Conf.LoadConf("user", "token")
	if username == "" || token == "" {

	} else {
		url := "/panel"
		http.Redirect(w, r, url, http.StatusFound)
	}
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("html/login.html")
		if err != nil {
			fmt.Fprintf(w, "parse template error: %s", err.Error())
			return
		}
		t.Execute(w, nil)
	} else {
		username := r.Form["username"]
		password := r.Form["password"]
		fmt.Fprintf(w, "username = %s, password = %s", username, password)
	}
}

func panel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	username := Conf.LoadConf("user", "username")
	token := Conf.LoadConf("user", "token")
	if username == "" || token == "" {
		url := "/login"
		http.Redirect(w, r, url, http.StatusFound)
	}
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("html/panel.html")
		if err != nil {
			fmt.Fprintf(w, "parse template error: %s", err.Error())
			return
		}
		t.Execute(w, nil)
	} else {
		username := r.Form["username"]
		password := r.Form["password"]
		fmt.Fprintf(w, "username = %s, password = %s", username, password)
	}
}

func setting_get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	username := Conf.LoadConf("user", "username")
	token := Conf.LoadConf("user", "token")
	if username == "" || token == "" {
		url := "/login"
		http.Redirect(w, r, url, http.StatusFound)
	}
	_, ret, err := Net.Post("http://go.bilihp.com:180/v1/pc/setting/setting_get", map[string]interface{}{"username": username, "token": token}, nil, nil)
	//fmt.Println(ret.(string))
	if err != nil {
		fmt.Println("setting_get", err)
	} else {
		jsr, err := Jsong.JObject(ret.(string))
		if err != nil {
			fmt.Println("setting_get", err)
		} else {
			jsp, _ := Jsong.ParseObject(jsr["data"])
			for k, v := range jsp {
				Conf.SaveConf("setting", Calc.Any2String(k), Calc.Any2String(v))
			}
		}

		w.Write([]byte(ret.(string)))
	}
}

func setting_set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	username := Conf.LoadConf("user", "username")
	token := Conf.LoadConf("user", "token")
	if username == "" || token == "" {
		url := "/login"
		http.Redirect(w, r, url, http.StatusFound)
	}
	key := r.PostFormValue("key")
	value := r.PostFormValue("value")
	Conf.SaveConf("setting", Calc.Any2String(key), Calc.Any2String(value))
	_, ret, err := Net.Post("http://go.bilihp.com:180/v1/pc/setting/setting_set", map[string]interface{}{"username": username, "token": token, "key": key, "value": value}, nil, nil)
	//fmt.Println(ret.(string))
	if err != nil {
		fmt.Println("setting_set", err)
	} else {
		w.Write([]byte(ret.(string)))
	}

}
