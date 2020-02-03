package Calc

import (
	"strconv"
	"time"
)

func GenerateToken() string {
	unix := time.Now().UnixNano()
	rand := Rand(0, 99999999)
	str := strconv.FormatInt(unix, 10) + strconv.FormatInt(int64(rand), 10)
	return Md5(str)
}
