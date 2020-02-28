package ActionRoute

import (
	"fmt"
	"main.go/Conf"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"net"
)

func PCRoute(jobject map[string]interface{}, username string, conn *net.TCPConn) {
	route := jobject["route"]
	data := jobject["data"]
	echo := jobject["echo"]
	switch route {

	case "update_config":
		//fmt.Println(route, data, echo)
		jsp, err := Jsong.ParseObject(data)
		if err != nil {
			fmt.Println("设置更新失败：", err)
		} else {
			for k, v := range jsp {
				//fmt.Println(reflect.TypeOf(v), v)
				if v == false || Calc.Any2String(v) == "0" {
					Conf.SaveConf("setting", Calc.Any2String(k), "0")
				} else if v == true || Calc.Any2String(v) == "1" {
					Conf.SaveConf("setting", Calc.Any2String(k), "1")
				} else {
					Conf.SaveConf("setting", Calc.Any2String(k), Calc.Any2String(v))
				}
			}
			fmt.Println("设置实时更新完毕")
		}
		break

	default:
		fmt.Println(route, data, echo)
		break
	}
}
