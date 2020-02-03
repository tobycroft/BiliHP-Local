package Calc

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math/big"
	rand2 "math/rand"
	"time"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Sha1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}

func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func Sha512(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(str)))
}

func Mt_rand(min, max int64) int64 {
	rand2.Seed(Seed())
	if min == max {
		return min
	} else {
		r := rand2.New(rand2.NewSource(time.Now().UnixNano()))
		return r.Int63n(max-min+1) + min
	}
}

func Seed() int64 {
	num, _ := rand.Int(rand.Reader, big.NewInt(999999999))
	return num.Int64()
}

func Rand(min, max int) int {
	rand2.Seed(Seed())
	if min == max {
		return min
	} else {
		var randNum int
		if max-min < 0 {
			randNum = rand2.Intn(min-max) + min
		} else {
			randNum = rand2.Intn(max-min) + min
		}
		return randNum
	}
}

func Any2Int64(any interface{}) int64 {
	ret, err := String2Int64(Any2String(any))
	if err != nil {
		return 99999998
	}
	return ret
}

func Any2Float64(any interface{}) float64 {
	ret, err := String2Float64(Any2String(any))
	if err != nil {
		return 99999998
	}
	return ret
}

func Any2Int(any interface{}) int {
	ret, err := String2Int(Any2String(any))
	if err != nil {
		return 99999998
	}
	return ret
}
