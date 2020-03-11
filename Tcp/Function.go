package Tcp

import (
	"main.go/Action/ActionRoute"
	"main.go/Conf"
	"time"
)

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
		time.Sleep(time.Second * 59)
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
		time.Sleep(time.Second * 57)
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
