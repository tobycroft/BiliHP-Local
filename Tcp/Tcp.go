package Tcp

import (
	"fmt"
	"main.go/Action/ActionRoute"
	"main.go/tuuz/RET"
	"net"
	"os"
	"time"
)

const addr = "go.bilihp.com:181"

var Conn = make(map[string]*net.TCPConn)

func Create(username string, token string) {
	server := addr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: ", err)
		os.Exit(1)
	}

	//建立服务器连接
	Conn[username], err = net.DialTCP("tcp", nil, tcpAddr)
	data := make(map[string]interface{})
	data["username"] = username
	data["token"] = token
	Sender(username, RET.Ws_succ("init", 0, data, "init"))
	if err != nil {
		fmt.Println("连接故障……正在重连……")
		time.Sleep(time.Second)
		Create(username, token)
	} else {
		fmt.Println("成功连入服务器！")
		Handler(username)
	}

}

func Sender(username string, message string) {
	conn := Conn[username]
	words := message
	_, err := conn.Write([]byte(words)) //给服务器发信息

	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "服务器反馈")
		os.Exit(1)
	}
}

func Handler(username string) {
	conn := Conn[username]
	var temp string
	for {
		buf := make([]byte, 1460)
		n, _ := conn.Read(buf)
		if n == 1460 {
			temp += string(buf[:n])
		} else {
			temp += string(buf[:n])
			msg := temp
			temp = ""
			ActionRoute.ActionRoute(msg, username, conn)
		}
	}
}
