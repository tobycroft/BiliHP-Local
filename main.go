package main

import (
	"main.go/Conf"
	"main.go/Tcp"
)

func main() {
	username := Conf.LoadConf("user", "username")
	if username != "" {
		Tcp.Create(username)
	} else {

	}

}
