package config

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

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
	f, err := os.Open("./config.txt")
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
