package ActionRoute

import (
	"fmt"
	"net"
)

func ActionRoute(json string, username string, conn *net.TCPConn) {
	fmt.Println(json, username)

}
