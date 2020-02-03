package tuuz

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz/Preg"
	"main.go/tuuz/database"
	"runtime"
)

func Db() gorose.IOrm {
	return database.Database.NewOrm()
}

func FUNCTION() string {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	Preg.MatchOwn("[A-z]+$", &name)
	return name
}

func FUNCTION_ALL() string {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	return name
}
