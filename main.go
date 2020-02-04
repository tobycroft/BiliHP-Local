package main

import (
	"fmt"
	"html/template"
	"main.go/Conf"
	"main.go/Tcp"
	"net/http"
	"os/exec"
)

func main() {
	username := Conf.LoadConf("user", "username")
	token := Conf.LoadConf("user", "token")

	fmt.Println(token)
	if username != "" && token != "" {
		Tcp.Create(username, token)
	} else {
		exec.Command(`cmd`, `/c`, `start`, `http://localhost/`).Start()
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/writelogin", writelogin)
	http.HandleFunc("/user/login", UserLoginHandler)
	// 设置静态目录
	fsh := http.FileServer(http.Dir("./html"))
	http.Handle("/html/", http.StripPrefix("/html/", fsh))

	err := http.ListenAndServe("0.0.0.0:80", nil)

	if err != nil {
		fmt.Println("服务器错误")
	}
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("r.Method = ", r.Method)
	//fmt.Println("r.URL = ", r.URL)
	//fmt.Println("r.Header = ", r.Header)
	//fmt.Println("r.Body = ", r.Body)
	//fmt.Fprintf(w, "HelloWorld!")
}

func index(w http.ResponseWriter, r *http.Request) {
	username := Conf.LoadConf("user", "username")
	token := Conf.LoadConf("user", "token")
	if username == "" || token == "" {
		url := "/login"
		http.Redirect(w, r, url, http.StatusFound)

	} else {

	}
}

func writelogin(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Handler Hello")
	username := request.PostFormValue("username")
	token := request.PostFormValue("token")
	Conf.SaveConf("user", "username", username)
	Conf.SaveConf("user", "token", token)
	fmt.Fprintf(response, "Login Success")
}

func UserLoginHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Handler Hello")
	fmt.Fprintf(response, "Login Success")
}

func login(w http.ResponseWriter, r *http.Request) {
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
