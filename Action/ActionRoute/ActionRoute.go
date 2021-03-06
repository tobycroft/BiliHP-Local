package ActionRoute

import (
	"fmt"
	"main.go/Common/update"
	"main.go/Conf"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/RET"
	"net"
	"os"
	"strings"
	"time"
)

func ActionRoute(jobject map[string]interface{}, username string, conn *net.TCPConn) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("----ActionRoute-recover----")
			fmt.Println(err)
		}
	}()
	code := Calc.Any2Int(jobject["code"])
	if code == -1 {
		ecam2(conn, "[登录信息]：", "登录信息错误！", "")
		Conf.SaveConf("user", "username", "")
		Conf.SaveConf("user", "token", "")
		os.Exit(1)
		return
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
		ecam2(conn, "", ret, "")
		break

	case "info":
		ecam2(conn, "", ret, "")
		break

	case "warning":
		ecam2(conn, "", ret, "")
		break

	case "error":
		ecam2(conn, "", ret, "")
		break

	case "update":
		ecam2(conn, "", echo, "")
		update.DoUpdate()
		break

	case "force_update":
		fmt.Println(echo)
		//if Conf.SystemType() == "windows" {
		//	exec.Command(`cmd`, `/c`, `start`, Calc.Any2String(ret)).Start()
		//}
		//if Conf.SystemType() == "linux" {
		//	fmt.Println("请执行:" + Calc.Any2String(ret))
		//	fmt.Println("------------------START------------------")
		//	fmt.Println("wget " + Calc.Any2String(ret))
		//	fmt.Println("------------------END------------------")
		//}
		update.DoUpdate()
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
			ecam2(conn, "[BiliHP-Debug]:", ret, "")
		}
		break

	case "other":
		ecam2(conn, "[BiliHP-Other]:", ret, "")
		break

	case "ecam":
		ecam2(conn, "[BiliHP-ECAM]:", ret, "")
		break

	case "alert":
		ecam2(conn, "[BiliHP-Alert]:", ret, "")
		break

	case "login":
		ecam2(conn, "[BiliHP-Login]:", ret, "")
		break

	case "loged":
		ecam2(conn, "[BiliHP-Loged]:", ret, "")
		break

	case "clear":
		fmt.Println(ret)
		break

	case "notam":
		ecam2(conn, "[BiliHP-NOTAM]:", ret, "")
		break

	case "system":
		ecam2(conn, "[BiliHP-系统消息]:", ret, "")
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
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err2, "")
				break
			}
			var values, err3 = Jsong.ParseObject(rets["values"])
			if err3 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err3, "")
				break
			}
			var cookie, err4 = Jsong.ParseObject(rets["cookie"])
			if err4 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err4, "")
				break
			}
			var url = Calc.Any2String(rets["url"])
			var method = Calc.Any2String(rets["method"])
			var route = Calc.Any2String(rets["route"])
			var typ = Calc.Any2String(rets["type"])
			var delay = Calc.Any2Float64(rets["delay"])
			ecam2(conn, "", echo, "")
			go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
		}
		break

	case "join_room":
		rets, err := Jsong.ParseObject(ret)
		if err != nil {
			fmt.Println("CURL信息不正确")
		} else {
			var header, err2 = Jsong.ParseObject(rets["header"])
			if err2 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err2, "")
				break
			}
			var values, err3 = Jsong.ParseObject(rets["values"])
			if err3 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err3, "")
				break
			}
			var cookie, err4 = Jsong.ParseObject(rets["cookie"])
			if err4 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err4, "")
				break
			}
			var url = Calc.Any2String(rets["url"])
			var method = Calc.Any2String(rets["method"])
			var route = Calc.Any2String(rets["route"])
			var typ = Calc.Any2String(rets["type"])
			var delay = Calc.Any2Float64(rets["delay"])
			ecam2(conn, "", echo, "")
			if Conf.LoadConf("setting", "join_room") == "1" {
				go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
			}
		}
		break

	case "gift":
		rets, err := Jsong.ParseObject(ret)
		if err != nil {
			fmt.Println("CURL信息不正确")
		} else {
			var header, err2 = Jsong.ParseObject(rets["header"])
			if err2 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err2, "")
				break
			}
			var values, err3 = Jsong.ParseObject(rets["values"])
			if err3 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err3, "")
				break
			}
			var cookie, err4 = Jsong.ParseObject(rets["cookie"])
			if err4 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err4, "")
				break
			}
			var url = Calc.Any2String(rets["url"])
			var method = Calc.Any2String(rets["method"])
			var route = Calc.Any2String(rets["route"])
			var typ = Calc.Any2String(rets["type"])
			var delay = Calc.Any2Float64(rets["delay"])
			ecam2(conn, "", echo, "")
			if Conf.LoadConf("setting", "raffle") == "1" {
				Time := Conf.LoadConf("setting", "time")
				ok, reason := Gift_check(Time)
				if !ok {
					ecam2(conn, "", "[BiliHP-Security]+"+reason, "")
					break
				}
				Percent := Conf.LoadConf("setting", "percent")
				ok2, reason2 := Gift_ratio(Calc.Any2Int(Percent))
				if !ok2 {
					ecam2(conn, "", "[BiliHP-Security]+"+reason2, "")
					break
				}
				go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
			} else {
				ecam2(conn, "", "小电视-领取被关闭", "")
				//fmt.Println("小电视-领取被关闭")
			}
		}
		break

	case "guard":
		rets, err := Jsong.ParseObject(ret)
		if err != nil {
			fmt.Println("CURL信息不正确")
		} else {
			var header, err2 = Jsong.ParseObject(rets["header"])
			if err2 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err2, "")
				break
			}
			var values, err3 = Jsong.ParseObject(rets["values"])
			if err3 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err3, "")
				break
			}
			var cookie, err4 = Jsong.ParseObject(rets["cookie"])
			if err4 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err4, "")
				break
			}
			var url = Calc.Any2String(rets["url"])
			var method = Calc.Any2String(rets["method"])
			var route = Calc.Any2String(rets["route"])
			var typ = Calc.Any2String(rets["type"])
			var delay = Calc.Any2Float64(rets["delay"])
			ecam2(conn, "", echo, "")
			if Conf.LoadConf("setting", "guard") == "1" {
				Time := Conf.LoadConf("setting", "time")
				ok, reason := Gift_check(Time)
				if !ok {
					ecam2(conn, "", "[BiliHP-Security]+"+reason, "")
					break
				}
				Percent := Conf.LoadConf("setting", "percent")
				ok2, reason2 := Gift_ratio(Calc.Any2Int(Percent))
				if !ok2 {
					ecam2(conn, "", "[BiliHP-Security]+"+reason2, "")
					break
				}
				go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
			} else {
				ecam2(conn, "", "总督-领取被关闭", "")
			}
		}
		break

	case "tianxuan":
		rets, err := Jsong.ParseObject(ret)
		if err != nil {
			fmt.Println("CURL信息不正确")
		} else {
			var header, err2 = Jsong.ParseObject(rets["header"])
			if err2 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err2, "")
				break
			}
			var values, err3 = Jsong.ParseObject(rets["values"])
			if err3 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err3, "")
				break
			}
			var cookie, err4 = Jsong.ParseObject(rets["cookie"])
			if err4 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err4, "")
				break
			}
			var url = Calc.Any2String(rets["url"])
			var method = Calc.Any2String(rets["method"])
			var route = Calc.Any2String(rets["route"])
			var typ = Calc.Any2String(rets["type"])
			var delay = Calc.Any2Float64(rets["delay"])
			ecam2(conn, "", echo, "")
			if Conf.LoadConf("setting", "tianxuan") == "1" {
				Time := Conf.LoadConf("setting", "time")
				ok, reason := Gift_check(Time)
				if !ok {
					ecam2(conn, "", "[BiliHP-Security]+"+reason, "")
					break
				}
				Percent := Conf.LoadConf("setting", "percent")
				ok2, reason2 := Gift_ratio(Calc.Any2Int(Percent))
				if !ok2 {
					ecam2(conn, "", "[BiliHP-Security]+"+reason2, "")
					break
				}
				bw := Conf.LoadConf("setting", "ban_words")
				bws := strings.Split(bw, ",")
				bdm := Conf.LoadConf("setting", "ban_danmu")
				bdms := strings.Split(bdm, ",")
				wdm := Conf.LoadConf("setting", "white_words")
				wdms := strings.Split(wdm, ",")
				mr := Conf.LoadConf("setting", "medal_room")
				mrs := strings.Split(mr, ",")
				br := Conf.LoadConf("setting", "ban_room")
				brs := strings.Split(br, ",")
				obj, ess := Jsong.ParseObject(jobject["object"])
				cont := true
				if ess != nil {

				} else {
					for _, word := range brs {
						if Calc.Any2String(obj["room_id"]) == word && len(word) > 1 {
							cont = false
							fmt.Println("触发天选屏蔽房间：", bws)
							ecam2(conn, "", "天选时刻-房间号"+Calc.Any2String(obj["room_id"])+"在屏蔽房间("+word+")中，不参与", "")
							break
						}
					}
					if cont {
						for _, word := range bws {
							if strings.Contains(strings.ToLower(Calc.Any2String(obj["award_name"])), strings.ToLower(word)) && len(word) > 1 {
								cont = false
								fmt.Println("设定屏蔽词：", bws)
								ecam2(conn, "", "天选时刻-奖品"+Calc.Any2String(obj["award_name"])+"与("+word+")匹配，不参与", "")
								break
							}
						}
					}
					if cont {
						for _, dm := range bdms {
							if strings.Contains(strings.ToLower(Calc.Any2String(obj["danmu"])), strings.ToLower(dm)) && len(dm) > 1 {
								cont = false
								fmt.Println("触发弹幕屏蔽词：", dm)
								ecam2(conn, "", "天选时刻-弹幕"+Calc.Any2String(obj["danmu"])+"与("+dm+")匹配，不参与", "")
								break
							}
						}
					}
					if Conf.LoadConf("setting", "blacklist_first") == "1" {
						if !cont {
							break
						}
					}
					if Conf.LoadConf("setting", "use_white") == "1" {
						for _, word := range wdms {
							cont = false
							if strings.Contains(strings.ToLower(Calc.Any2String(obj["award_name"])), strings.ToLower(word)) && len(word) > 1 {
								cont = true
								fmt.Println("触发天选白名单：", bws)
								ecam2(conn, "", "天选时刻-奖品"+Calc.Any2String(obj["award_name"])+"与("+word+")匹配，正常参与", "")
								break
							}
						}
					}
					if cont && Calc.Any2String(obj["need_medal"]) == "1" {
						cont = false
						for _, word := range mrs {
							if Calc.Any2String(obj["room_id"]) == word && len(word) > 1 {
								cont = true
								break
							}
							if !cont {
								ecam2(conn, "", "你没有该主播抽奖所需的需要2级勋章，自动跳过", "")
							}
						}
					}
				}
				if cont {
					go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
				}
			} else {
				ecam2(conn, "", "天选时刻-领取被关闭", "")
			}
		}
		break

	case "box":
		rets, err := Jsong.ParseObject(ret)
		if err != nil {
			fmt.Println("CURL信息不正确")
		} else {
			var header, err2 = Jsong.ParseObject(rets["header"])
			if err2 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err2, "")
				break
			}
			var values, err3 = Jsong.ParseObject(rets["values"])
			if err3 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err3, "")
				break
			}
			var cookie, err4 = Jsong.ParseObject(rets["cookie"])
			if err4 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err4, "")
				break
			}
			var url = Calc.Any2String(rets["url"])
			var method = Calc.Any2String(rets["method"])
			var route = Calc.Any2String(rets["route"])
			var typ = Calc.Any2String(rets["type"])
			var delay = Calc.Any2Float64(rets["delay"])

			ecam2(conn, "", echo, "")

			if Conf.LoadConf("setting", "box") == "1" {
				Time := Conf.LoadConf("setting", "time")
				ok, reason := Gift_check(Time)
				if !ok {
					ecam2(conn, "", "[BiliHP-Security]+"+reason, "")
					break
				}
				Percent := Conf.LoadConf("setting", "percent")
				ok2, reason2 := Gift_ratio(Calc.Any2Int(Percent))
				if !ok2 {
					ecam2(conn, "", "[BiliHP-Security]+"+reason2, "")
					break
				}
				bw := Conf.LoadConf("setting", "ban_words")
				bws := strings.Split(bw, ",")
				obj, ess := Jsong.ParseObject(jobject["object"])
				cont := true
				if ess != nil {

				} else {
					for _, word := range bws {
						if strings.Contains(Calc.Any2String(obj["title"]), word) && len(word) > 1 {
							cont = false
							fmt.Println("设定屏蔽词：", bws)
							ecam2(conn, "", "活动抽奖-奖品"+Calc.Any2String(obj["award_name"])+"与("+word+")匹配，不参与", "")
							break
						}
					}
				}
				if cont {
					go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
				}

			} else {
				ecam2(conn, "", "活动抽奖-领取被关闭", "")
			}
		}
		break

	case "pk":
		rets, err := Jsong.ParseObject(ret)
		if err != nil {
			fmt.Println("CURL信息不正确")
		} else {
			var header, err2 = Jsong.ParseObject(rets["header"])
			if err2 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err2, "")
				break
			}
			var values, err3 = Jsong.ParseObject(rets["values"])
			if err3 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err3, "")
				break
			}
			var cookie, err4 = Jsong.ParseObject(rets["cookie"])
			if err4 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err4, "")
				break
			}
			var url = Calc.Any2String(rets["url"])
			var method = Calc.Any2String(rets["method"])
			var route = Calc.Any2String(rets["route"])
			var typ = Calc.Any2String(rets["type"])
			var delay = Calc.Any2Float64(rets["delay"])
			ecam2(conn, "", echo, "")
			if Conf.LoadConf("setting", "pk") == "1" {
				Time := Conf.LoadConf("setting", "time")
				ok, reason := Gift_check(Time)
				if !ok {
					ecam2(conn, "", "[BiliHP-Security]+"+reason, "")
					break
				}
				Percent := Conf.LoadConf("setting", "percent")
				ok2, reason2 := Gift_ratio(Calc.Any2Int(Percent))
				if !ok2 {
					ecam2(conn, "", "[BiliHP-Security]+"+reason2, "")
					break
				}
				go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
			} else {
				ecam2(conn, "", "天选时刻-领取被关闭", "")
			}
		}
		break

	case "storm", "STORM":
		rets, err := Jsong.ParseObject(ret)
		if err != nil {
			fmt.Println("CURL信息不正确")
		} else {
			var header, err2 = Jsong.ParseObject(rets["header"])
			if err2 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err2, "")
				break
			}
			var values, err3 = Jsong.ParseObject(rets["values"])
			if err3 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err3, "")
				break
			}
			var cookie, err4 = Jsong.ParseObject(rets["cookie"])
			if err4 != nil {
				ecam2(conn, "[BiliHP-LOCAL-ERROR]:", err4, "")
				break
			}
			var url = Calc.Any2String(rets["url"])
			var method = Calc.Any2String(rets["method"])
			var route = Calc.Any2String(rets["route"])
			var typ = Calc.Any2String(rets["type"])
			var delay = Calc.Any2Float64(rets["delay"])
			ecam2(conn, "", echo, "")
			if Conf.LoadConf("setting", "storm") == "1" {
				Time := Conf.LoadConf("setting", "time")
				ok, reason := Gift_check(Time)
				if !ok {
					ecam2(conn, "", "[BiliHP-Security]+"+reason, "")
					break
				}
				Percent := Conf.LoadConf("setting", "percent")
				ok2, reason2 := Gift_ratio(Calc.Any2Int(Percent))
				if !ok2 {
					ecam2(conn, "", "[BiliHP-Security]+"+reason2, "")
					break
				}
				go Curl(url, method, values, header, cookie, typ, echo, *conn, route, delay)
			} else {
				ecam2(conn, "", "节奏风暴-领取被关闭", "")
			}
		}
		break

	default:
		fmt.Println("undefine-route", typ, ret, echo)
		break

	}
}

func Gift_check(Time string) (bool, string) {
	timing := time.Now().Hour()
	times, err := Jsong.JObject(Time)
	if err != nil {
		return true, "解析出错，本时段自动启用"
	} else {
		for timer, bool := range times {
			if timer == "t"+Calc.Any2String(timing) {
				if bool == true {
					return true, "本时段可用"
				}
			}

		}
	}
	return false, "本时段用户不参与抢礼物，如需启用，请在PC/C2C远程设置中开启"
}

func Gift_ratio(ratio int) (bool, string) {
	num := Calc.Rand(1, 100)
	if num < ratio {
		return true, "概率系统自动捕捉" + Calc.Any2String(num)
	} else {
		return false, "概率系统自动跳过" + Calc.Any2String(num)
	}
}

func ecam(msg interface{}, ret interface{}, color string) {
	fmt.Println(msg, ret, color)
}

func ecam2(conn *net.TCPConn, msg interface{}, ret interface{}, color string) {
	ecam(msg, ret, color)
	Send(*conn, SendObj("send_app", msg, "", ret))
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
	server := Conf.Addr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Println("重连故障：", err)
	}
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
}
