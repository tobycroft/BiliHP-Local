package ActionRoute

import (
	"fmt"
	"net"
)

func PCRoute(jobject map[string]interface{}, username string, conn *net.TCPConn) {
	route := jobject["route"]
	data := jobject["data"]
	echo := jobject["echo"]
	switch route {

	case "update_config":
		fmt.Println(route, data, echo)
		break

	default:
		fmt.Println(route, data, echo)
		break
	}
}
