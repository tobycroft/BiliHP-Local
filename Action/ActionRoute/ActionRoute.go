package ActionRoute

import (
	"fmt"
	"main.go/Conf"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/RET"
	"net"
	"os"
	"os/exec"
	"time"
)

const Addr = "go.bilihp.com:181"

func ActionRoute(json string, username string, conn *net.TCPConn) {
	jsons, err := Jsong.TCPJObject(json)
	if err != nil {
		fmt.Println("err:", err, json)
	} else {
		for _, jobject := range jsons {
			code := Calc.Any2Int(jobject["code"])
			if code == -1 {
				ecam("[登录信息]：", "登录信息错误！", "")
				Conf.SaveConf("user", "username", "")
				Conf.SaveConf("user", "token", "")
				os.Exit(1)
				break
			}
			typ := Calc.Any2String(jobject["type"])
			ret := jobject["data"]
			echo := Calc.Any2String(jobject["echo"])

			switch typ {
			case "orign":
				fmt.Println(ret)
				break

			case "app":
				PCRoute(jobject, username, conn)
				break

			case "supercurl":
				fmt.Println(ret)
				break

			case "info":
				fmt.Println(ret)
				break

			case "warning":
				fmt.Println(ret)
				break

			case "error":
				fmt.Println(ret)
				break

			case "update":
				fmt.Println(echo)
				exec.Command(`cmd`, `/c`, `start`, Calc.Any2String(ret)).Start()
				time.Sleep(120 * time.Second)
				os.Exit(1)
				break

			case "reinit":
				token := Conf.LoadConf("user", "token")
				data := make(map[string]interface{})
				data["username"] = username
				data["token"] = token
				Send(*conn, RET.Ws_succ("init", 0, data, "init"))
				break

			case "debug":
				if Conf.LoadConf("debug", "debug") == "true" {
					ecam("[BiliHP-Debug]:", ret, "")
				}
				break

			case "other":
				ecam("[BiliHP-Other]:", ret, "")
				break

			case "ecam":
				ecam("[BiliHP-ECAM]:", ret, "")
				break

			case "alert":
				ecam("[BiliHP-Alert]:", ret, "")
				break

			case "login":
				ecam("[BiliHP-Login]:", ret, "")
				break

			case "loged":
				ecam("[BiliHP-Loged]:", ret, "")
				break

			case "clear":
				fmt.Println(ret)
				break

			case "notam":
				ecam("[BiliHP-NOTAM]:", ret, "")
				break

			case "system":
				ecam("[BiliHP-系统消息]:", ret, "")
				break

			case "pong":
				if Conf.LoadConf("debug", "debug") == "true" {
					ecam("[BiliHP-Ping]:", ret, "")
				}
				break

			case "curl":
				rets, err := Jsong.ParseObject(ret)
				if err != nil {
					fmt.Println("CURL信息不正确")
				} else {
					var header, err2 = Jsong.ParseObject(rets["header"])
					if err2 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err2, "")
						break
					}
					var values, err3 = Jsong.ParseObject(rets["values"])
					if err3 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err3, "")
						break
					}
					var cookie, err4 = Jsong.ParseObject(rets["cookie"])
					if err4 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err4, "")
						break
					}
					var url = Calc.Any2String(rets["url"])
					var method = Calc.Any2String(rets["method"])
					var route = Calc.Any2String(rets["route"])
					var typ = Calc.Any2String(rets["type"])
					var delay = Calc.Any2Float64(rets["delay"])
					ecam("", echo, "")
					go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
				}
				break

			case "gift":
				rets, err := Jsong.ParseObject(ret)
				if err != nil {
					fmt.Println("CURL信息不正确")
				} else {
					var header, err2 = Jsong.ParseObject(rets["header"])
					if err2 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err2, "")
						break
					}
					var values, err3 = Jsong.ParseObject(rets["values"])
					if err3 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err3, "")
						break
					}
					var cookie, err4 = Jsong.ParseObject(rets["cookie"])
					if err4 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err4, "")
						break
					}
					var url = Calc.Any2String(rets["url"])
					var method = Calc.Any2String(rets["method"])
					var route = Calc.Any2String(rets["route"])
					var typ = Calc.Any2String(rets["type"])
					var delay = Calc.Any2Float64(rets["delay"])
					ecam("", echo, "")
					go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
				}
				break

			case "guard":
				rets, err := Jsong.ParseObject(ret)
				if err != nil {
					fmt.Println("CURL信息不正确")
				} else {
					var header, err2 = Jsong.ParseObject(rets["header"])
					if err2 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err2, "")
						break
					}
					var values, err3 = Jsong.ParseObject(rets["values"])
					if err3 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err3, "")
						break
					}
					var cookie, err4 = Jsong.ParseObject(rets["cookie"])
					if err4 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err4, "")
						break
					}
					var url = Calc.Any2String(rets["url"])
					var method = Calc.Any2String(rets["method"])
					var route = Calc.Any2String(rets["route"])
					var typ = Calc.Any2String(rets["type"])
					var delay = Calc.Any2Float64(rets["delay"])
					ecam("", echo, "")
					go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
				}
				break

			case "tianxuan":
				rets, err := Jsong.ParseObject(ret)
				if err != nil {
					fmt.Println("CURL信息不正确")
				} else {
					var header, err2 = Jsong.ParseObject(rets["header"])
					if err2 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err2, "")
						break
					}
					var values, err3 = Jsong.ParseObject(rets["values"])
					if err3 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err3, "")
						break
					}
					var cookie, err4 = Jsong.ParseObject(rets["cookie"])
					if err4 != nil {
						ecam("[BiliHP-LOCAL-ERROR]:", err4, "")
						break
					}
					var url = Calc.Any2String(rets["url"])
					var method = Calc.Any2String(rets["method"])
					var route = Calc.Any2String(rets["route"])
					var typ = Calc.Any2String(rets["type"])
					var delay = Calc.Any2Float64(rets["delay"])
					ecam("", echo, "")
					go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
				}
				break

			default:
				fmt.Println("undefine-route", typ, ret, echo)
				break

			}
		}
	}
}

func ecam(msg interface{}, ret interface{}, color string) {
	fmt.Println(msg, ret, color)
}

func SendObj(typ string, data interface{}, echo string, values interface{}) string {
	obj := make(map[string]interface{})
	obj["type"] = typ
	obj["data"] = data
	obj["echo"] = echo
	obj["values"] = values
	ret, _ := Jsong.Encode(obj)
	return ret
}

func Send(conn net.TCPConn, message string) bool {
	words := message
	//fmt.Println(words)
	_, err := conn.Write([]byte(words)) //给服务器发信息

	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "服务器发送失败，检测到断线，开始重连")
		return false
	} else {
		return true
	}
}

func Reconnect(conn *net.TCPConn) {
	server := Addr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Println("重连故障：", err)
	}
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
}
