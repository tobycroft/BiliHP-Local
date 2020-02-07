package Tcp

import (
	"fmt"
	"main.go/Action/ActionRoute"
	"main.go/tuuz/RET"
	"net"
	"os"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var Conn = make(map[string]*net.TCPConn)

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
	} else {
		fmt.Println("成功连入服务器！")
		data := make(map[string]interface{})
		data["username"] = username
		data["token"] = token
		data["version"] = "0.14.0"
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
	defer wg.Done()
	conn := Conn[username]
	go update_setting()
	go yingyuan_sign(*conn)
	go daily_task(*conn)
	go silver_task(*conn)
	go online_silver(*conn)
	go daily_bag(*conn)
	go app_heart(*conn)
	go pc_heart(*conn)
	go ping(*conn)
	wg.Wait()
}

func Handler(username string, token string) {
	conn := Conn[username]
	var temp string
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			wg.Done()
			fmt.Println("handler出错:", err)
		}
		//fmt.Println("len:", n, err)
		if n >= 1024 {
			temp += string(buf[:n])
		} else {
			temp += string(buf[:n])
			msg := temp
			temp = ""
			//fmt.Println(msg)
			ActionRoute.ActionRoute(msg, username, conn)
		}
	}
}
