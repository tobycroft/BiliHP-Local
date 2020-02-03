package main

import (
	"fmt"
	"main.go/Conf"
	"main.go/Tcp"
	"net/http"
)

func main() {
	username := Conf.LoadConf("user", "username")
	token := Conf.LoadConf("user", "token")
	fmt.Println(token)
	if username != "" {
		Tcp.Create(username, token)
	} else {

	}
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)
	fmt.Fprintf(w, "HelloWorld!")
}

func UserLoginHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Handler Hello")
	fmt.Fprintf(response, "Login Success")
}
