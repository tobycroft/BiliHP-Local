package Calc

import "strconv"

func Int2String(num int) string {
	return strconv.Itoa(num)
}

func Int642String(num int64) string {
	return strconv.FormatInt(num, 10)
}

func Float642String(f64 float64) string {
	return strconv.FormatFloat(f64, 'f', -1, 64)
}
