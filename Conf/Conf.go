package Conf

import (
	"github.com/Unknwon/goconfig"
)

func LoadConf(section string, key string) string {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		return ""
	}

	value, err := cfg.GetValue("user", "username")
	if err != nil {
		return ""
	} else {
		return value
	}
}

func SaveConf(section string, key string, value string) bool {
	cfg, err := goconfig.LoadConfigFile("../conf.ini")
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
