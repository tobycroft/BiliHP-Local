package main

import (
	"fmt"
	"main.go/Conf"
)

func main() {
	fmt.Println(Conf.LoadConf("user", "username"))

}
