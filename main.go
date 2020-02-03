package main

import (
	"main.go/Conf"
	"main.go/Tcp"
)

func main() {
	username := Conf.LoadConf("user", "username")
	token := Conf.LoadConf("user", "token")
	if username != "" {
		Tcp.Create(username, token)
	} else {

	}

}
