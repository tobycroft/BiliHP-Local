package main

import (
	"fmt"
	"main.go/Conf"
	"main.go/Tcp"
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
