package main

import (
	"main.go/Conf"
	"main.go/Tcp"
)

func main() {
	username := Conf.LoadConf("user", "username")
	token := "4e26d35b2a93c7890fa3cf16d4e8b456"
	if username != "" {
		Tcp.Create(username, token)
	} else {

	}

}
