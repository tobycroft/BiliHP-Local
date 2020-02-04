package Tcp

import (
	"main.go/Action/ActionRoute"
	"main.go/Conf"
	"net"
	"time"
)

func yingyuan_sign(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "yingyuan_sign") == "1" {
			ret := ActionRoute.SendObj("func", nil, "yingyuan_sign", nil)
			ActionRoute.Send(conn, ret)
		}
		time.Sleep(time.Hour * 24)
	}

}

func daily_task(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "daily_task", nil)
			ActionRoute.Send(conn, ret)
		}
		time.Sleep(time.Hour * 24)
	}
}

func silver_task(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "silver_task", nil)
			ActionRoute.Send(conn, ret)
		}
		time.Sleep(time.Minute * 10)
	}

}

func online_silver(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "online_silver", nil)
			ActionRoute.Send(conn, ret)
		}
		time.Sleep(time.Minute * 10)
	}
}

func daily_bag(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "daily_bag") == "1" {
			ret := ActionRoute.SendObj("func", nil, "daily_bag", nil)
			ActionRoute.Send(conn, ret)
		}
		time.Sleep(time.Hour * 24)
	}
}

func app_heart(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "app_heart") == "1" {
			ret := ActionRoute.SendObj("func", nil, "app_heart", nil)
			ActionRoute.Send(conn, ret)
		}
		time.Sleep(time.Second * 30)
	}
}

func pc_heart(conn net.TCPConn) {
	for {
		if Conf.LoadConf("setting", "pc_heart") == "1" {
			ret := ActionRoute.SendObj("func", nil, "pc_heart", nil)
			ActionRoute.Send(conn, ret)
		}
		time.Sleep(time.Second * 30)
	}
}

func ping(conn net.TCPConn) {
	for {
		ret := ActionRoute.SendObj("ping", "ping", "ping", nil)
		ActionRoute.Send(conn, ret)
		time.Sleep(time.Second * 10)
	}
}
