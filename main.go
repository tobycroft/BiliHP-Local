package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		panic("错误")
	}

	value, err := cfg.GetValue("user", "username")
	fmt.Println(value)
	cfg.SetValue("user", "username", "dda")
	goconfig.SaveConfigFile(cfg, "conf.ini")
}
