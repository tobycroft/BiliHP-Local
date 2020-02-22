package Tcp

import (
	"fmt"
	"main.go/Action/ActionRoute"
	"main.go/Conf"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Net"
	"time"
)

func update_setting() {
	for {
		username := Conf.LoadConf("user", "username")
		token := Conf.LoadConf("user", "token")
		if username == "" || token == "" {

		} else {
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

			}
		}
		time.Sleep(time.Minute * 1)
	}
}

func yingyuan_sign(username string) {

	for {
		if Conf.LoadConf("setting", "yingyuan_sign") == "1" {
			ret := ActionRoute.SendObj("func", nil, "yingyuan_sign", nil)
			Lock.Lock()
			if Conn[username] != nil {
				go ActionRoute.Send(*Conn[username], ret)
			}
			Lock.Unlock()
		}
		time.Sleep(time.Hour * 24)
	}
}

func daily_task(username string) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "daily_task", nil)
			Lock.Lock()
			if Conn[username] != nil {
				go ActionRoute.Send(*Conn[username], ret)
			}
			Lock.Unlock()
		}
		time.Sleep(time.Hour * 24)
	}

}

func silver_task(username string) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "silver_task", nil)
			Lock.Lock()
			if Conn[username] != nil {
				go ActionRoute.Send(*Conn[username], ret)
			}
			Lock.Unlock()
		}
		time.Sleep(time.Minute * 10)
	}

}

func online_silver(username string) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "online_silver", nil)
			Lock.Lock()
			if Conn[username] != nil {
				go ActionRoute.Send(*Conn[username], ret)
			}
			Lock.Unlock()
		}
		time.Sleep(time.Minute * 10)
	}

}

func daily_bag(username string) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "daily_bag", nil)
			Lock.Lock()
			if Conn[username] != nil {
				go ActionRoute.Send(*Conn[username], ret)
			}
			Lock.Unlock()
		}
		time.Sleep(time.Hour * 24)
	}

}

func app_heart(username string) {
	for {
		if Conf.LoadConf("setting", "app_heart") == "1" {
			ret := ActionRoute.SendObj("func", nil, "app_heart", nil)
			Lock.Lock()
			if Conn[username] != nil {
				go ActionRoute.Send(*Conn[username], ret)
			}
			Lock.Unlock()
		}
		time.Sleep(time.Second * 30)
	}
}

func pc_heart(username string) {
	for {
		if Conf.LoadConf("setting", "pc_heart") == "1" {
			ret := ActionRoute.SendObj("func", nil, "pc_heart", nil)
			Lock.Lock()
			if Conn[username] != nil {
				go ActionRoute.Send(*Conn[username], ret)
			}
			Lock.Unlock()
		}
		time.Sleep(time.Second * 30)
	}
}

func Ping(username string) {
	for {
		ret := ActionRoute.SendObj("ping", "ping", "ping", nil)
		Lock.Lock()
		if Conn[username] != nil {
			go ActionRoute.Send(*Conn[username], ret)
		}
		Lock.Unlock()
		time.Sleep(time.Second * 10)
	}
}
