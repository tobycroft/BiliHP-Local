package Tcp

import (
	"fmt"
	"main.go/Action/ActionRoute"
	"main.go/Conf"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Net"
	"net"
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

func yingyuan_sign(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "yingyuan_sign") == "1" {
			ret := ActionRoute.SendObj("func", nil, "yingyuan_sign", nil)
			if ActionRoute.Send(conn, ret) != true {
				wg.Done()
			}
		}
		time.Sleep(time.Hour * 24)
	}
}

func daily_task(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "daily_task", nil)
			if ActionRoute.Send(conn, ret) != true {
				wg.Done()
			}
		}
		time.Sleep(time.Hour * 24)
	}

}

func silver_task(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "silver_task", nil)
			if ActionRoute.Send(conn, ret) != true {
				wg.Done()
			}
		}
		time.Sleep(time.Minute * 10)
	}

}

func online_silver(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "online_silver", nil)
			if ActionRoute.Send(conn, ret) != true {
				wg.Done()
			}
		}
		time.Sleep(time.Minute * 10)
	}

}

func daily_bag(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "daily_bag", nil)
			if ActionRoute.Send(conn, ret) != true {
				wg.Done()
			}
		}
		time.Sleep(time.Hour * 24)
	}

}

func app_heart(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "app_heart") == "1" {
			ret := ActionRoute.SendObj("func", nil, "app_heart", nil)
			if ActionRoute.Send(conn, ret) != true {
				wg.Done()
			}
		}
		time.Sleep(time.Second * 30)
	}
}

func pc_heart(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "pc_heart") == "1" {
			ret := ActionRoute.SendObj("func", nil, "pc_heart", nil)
			if ActionRoute.Send(conn, ret) != true {
				wg.Done()
			}
		}
		time.Sleep(time.Second * 30)
	}
}

func ping(conn net.TCPConn) {
	for {
		ret := ActionRoute.SendObj("ping", "ping", "ping", nil)
		if ActionRoute.Send(conn, ret) != true {
			wg.Done()
		}
		time.Sleep(time.Second * 10)
	}
}
