package Conf

import (
	"github.com/Unknwon/goconfig"
	"runtime"
)

const Version = "1.5.2"

const Addr = "go.bilihp.com:181"

//const Addr = "127.0.0.1:80"

func SystemType() string {
	sysType := runtime.GOOS
	return sysType
}

func LoadConf(section string, key string) string {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		SaveConf(section, key, " ")
		return ""
	}
	value, err := cfg.GetValue(section, key)
	if err != nil {
		return ""
	} else {
		return value
	}
}

func LoadSec(section string) map[string]string {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	value, err := cfg.GetSection(section)
	if err != nil {
		return nil
	} else {
		return value
	}
}

func SaveConf(section string, key string, value string) bool {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		return false
	}
	cfg.SetValue(section, key, value)
	err = goconfig.SaveConfigFile(cfg, "conf.ini")
	if err != nil {
		return false
	} else {
		return true
	}
}
