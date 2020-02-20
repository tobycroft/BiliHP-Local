package Tcp

import (
	"fmt"
	"main.go/Action/ActionRoute"
	"main.go/Conf"
	"main.go/tuuz/RET"
	"net"
	"os"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var Conn = make(map[string]*net.TCPConn)
var MaxMSS = 0
var Lock sync.RWMutex

func Create(username string, token string) {
	defer func() {
		time.Sleep(1 * time.Second)
		Create(username, token)
	}()
	server := ActionRoute.Addr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: ", err)
		os.Exit(1)
	}

	//建立服务器连接
	Conn[username], err = net.DialTCP("tcp", nil, tcpAddr)
	wg.Add(1)
	if err != nil {
		fmt.Println("连接故障……正在重连……")

		time.Sleep(1 * time.Second)
		wg.Done()
	} else {
		fmt.Println("成功连入服务器！")
		data := make(map[string]interface{})
		data["username"] = username
		data["token"] = token
		data["version"] = Conf.Version
		data["type"] = "pc"
		Sender(username, token, RET.Ws_succ("init", 0, data, "init"))
		go Functions(username, token)
		Handler(username, token)
	}
	wg.Wait()
}

func Sender(username string, token string, message string) {
	conn := Conn[username]
	words := message
	_, err := conn.Write([]byte(words)) //给服务器发信息

	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "服务器发送失败正在重新建立重连……")
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}
}

func Functions(username string, token string) {
	go yingyuan_sign(username)
	go daily_task(username)
	go silver_task(username)
	go online_silver(username)
	go daily_bag(username)
	go app_heart(username)
	go pc_heart(username)
	go ping(username)
}

func Set_setting(username, key, value string) {
	data := make(map[string]string)
	data["key"] = key
	data["value"] = value

	rt := RET.Ws_succ2("app", "pc_set_setting", 0, data, "pc_set_setting")
	//str, _ := Jsong.Encode(data)
	if Conn[username] != nil {
		ActionRoute.Send(*Conn[username], rt)
	}
}

func Get_settings(username string) {
	data := make(map[string]string)

	rt := RET.Ws_succ2("app", "get_config", 0, data, "get_config")
	//str, _ := Jsong.Encode(data)
	if Conn[username] != nil {
		ActionRoute.Send(*Conn[username], rt)
	}
}

func Handler(username string, token string) {
	conn := Conn[username]
	var temp string
	var MaxMSS int = 0
	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			wg.Done()

			fmt.Println("handler出错:", err)
			return
		}
		//fmt.Println("len:", n, err)
		if n == MaxMSS {
			temp += string(buf[:n])
		} else {
			if n > MaxMSS {
				MaxMSS = n
			}
			temp += string(buf[:n])
			msg := temp
			temp = ""
			//fmt.Println(msg)
			ActionRoute.ActionRoute(msg, username, conn)
		}
	}
}
