package Tcp

import (
	"fmt"
	"main.go/Action/ActionRoute"
	"main.go/Conf"
	"main.go/tuuz/Jsong"
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
	server := Conf.Addr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: ", err)
		os.Exit(1)
	}
	//fmt.Println(tcpAddr)
	//建立服务器连接
	Conn[username], err = net.DialTCP("tcp", nil, tcpAddr)
	wg.Add(1)
	if err != nil {
		fmt.Println("连接故障……正在重连……", err)
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
		//fmt.Println("init")
		Functions(username, token)
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
	go do_sign(username)
	go manga_sign(username)
	go manga_share(username)
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
	temp := ""
	buf := make([]byte, 4096)
	for {
		n, err := Conn[username].Read(buf)
		if n == 0 || err != nil {
			wg.Done()
			fmt.Println("handler出错:", err)
			return

			//将当前用户从在线字典中删除
			//address := strings.Split(addr, ":")
			//ip := address[0]
			//Conclose(conn, addr, ip)
			//通知其他客户端该用户退出登录
			//Messages <- "logout"
			//Send_All("logout")
			return
		}
		temp += string(buf[:n])
		jsons, err := Jsong.TCPJObject(&temp)
		if err != nil {
			fmt.Println("err:", err)
		} else {
			for _, jobject := range jsons {
				go ActionRoute.ActionRoute(jobject, username, Conn[username])
			}
		}

	}

}
