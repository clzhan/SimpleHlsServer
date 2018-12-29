package conf

import (
	"fmt"
	"os"

	"runtime"
	"github.com/clzhan/SimpleHlsServer/utils"
)
var ostype = runtime.GOOS

var AppConf struct {
	WebPort       string
	IPlocal       string

}


func Init() {

	config := new(Config)
	var Ini string
	if ostype == "windows"{
		Ini = util.GetProjectPath() + "\\" + "mediaserver.ini"
	}else{
		Ini = util.GetProjectPath() + "/"+ "mediaserver.ini"
	}

	err := config.LoadConfig(Ini)

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}


	AppConf.IPlocal, err = config.ReadKeyValue("MediaSever","IP");
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	AppConf.WebPort, err = config.ReadKeyValue("MediaSever","webport");
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}


	fmt.Println("conf : ", AppConf)

}