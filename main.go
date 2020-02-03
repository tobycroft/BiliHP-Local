package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(GetConfig())
	//conn, err := net.Dial("tcp", "10.0.0.181:81")
	//defer conn.Close()
	//if err != nil {
	//	fmt.Printf("connect failed, err : %v\n", err.Error())
	//	return
	//}
	//
	//inputReader := bufio.NewReader(os.Stdin)
	//
	//for {
	//	input, err := inputReader.ReadString('\n')
	//	if err != nil {
	//		fmt.Printf("read from console failed, err: %v\n", err)
	//		break
	//	}
	//	trimmedInput := strings.TrimSpace(input)
	//	if trimmedInput == "Q" {
	//		break
	//	}
	//	_, err = conn.Write([]byte(trimmedInput))
	//
	//	if err != nil {
	//		fmt.Printf("write failed , err : %v\n", err)
	//		break
	//	}
	//}
}

type Config struct {
	Username string
	Token    string
}

func GetConfig() Config {
	//创建一个空的结构体,将本地文件读取的信息放入
	c := &Config{}
	//创建一个结构体变量的反射
	cr := reflect.ValueOf(c).Elem()
	//打开文件io流
	f, err := os.Open("./config.ini")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	//我们要逐行读取文件内容
	s := bufio.NewScanner(f)
	for s.Scan() {
		//以=分割,前面为key,后面为value
		var str = s.Text()
		var index = strings.Index(str, "=")
		var key = str[0:index]
		var value = str[index+1:]
		//因为Port是int,所以我们这里要将截取的string强转成int
		if strings.Contains(key, "Port") {
			var i, err = strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			//通过反射将字段设置进去
			cr.FieldByName(key).Set(reflect.ValueOf(i))
		} else {
			//通过反射将字段设置进去
			fmt.Println(key, value)
			cr.FieldByName(key).Set(reflect.ValueOf(value))
		}

	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	//返回Config结构体变量
	return *c
}
