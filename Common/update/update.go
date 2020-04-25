package update

import (
	"fmt"
	"github.com/inconshreveable/go-update"
	"net/http"
	"runtime"
)

func doUpdate() error {
	resp, err := http.Get("http://pandorabox.tuuz.cc:8000/app/" + version())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		fmt.Println("自动更新错误：", err)
		// error handling
	}
	return err
}

func version() string {
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			return "BiliHP_PCWEB.exe"

		default:
			return "BiliHP_PCWEB_386.exe"
		}

	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			return "c2c_linux"

		case "386":
			return "c2c_linux"

		case "mipsle":
			return "BiliHP_Router_linux"

		case "arm", "arm64":
			return "BiliHP_ARM_linux"

		case "mips":
			return "c2c_linux"

		default:
			return "c2c_linux"
		}

	case "darwin":
		return "BiliHP_Mac_darwin"

	default:
		fmt.Println("没有找到对应的版本")
		return "c2c_linux"
	}
}
