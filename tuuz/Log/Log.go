package Log

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger
var file *os.File

func Write(file_name string, logs string, discript string, other string) {
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("./log", os.ModePerm)
	}
	file, err := os.OpenFile("log/"+file_name+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		log.Fatalln(err)
	} else {
		logger := log.New(file, "", log.LstdFlags)
		logger.Println(logs, discript, other)
		file.Close()
	}
}

func Error(file_name string, err error) {
	Write(file_name, "", "", err)
}

func Err(err error) {
	Write("Error", "", "", err)
}

func Errs(err error, log string) {
	Write("Error", log, "", err)
}

//Database err
func Drr(err error) {
	Write("Database", "", "", err)
}

func Crr(logs interface{}) {
	Write("Common", "", "", logs)
}

func Crrs(logs interface{}, discript string) {
	Write("Common", "", discript, logs)
}

func Dbrr(err error, log string) {
	fmt.Println(err)
	Write("Dberror", log, "", err)
}
