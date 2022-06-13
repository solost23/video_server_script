package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"video_server_script/server"
)

var (
	WebConfigPath       = "config/config.yaml"
	version             = "__BUILD_VERSION__"
	execDir, scriptName string
	st, v, V            bool
)

func main() {
	flag.StringVar(&execDir, "d", ".", "脚本目录")
	flag.StringVar(&scriptName, "n", "import_video", "脚本名称")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本号")
	flag.BoolVar(&st, "s", false, "项目状态")
	flag.Parse()
	if v || V {
		fmt.Println(version)
		return
	}
	// 运行
	InitConfig()
	server.Run()
}

func InitConfig() {
	configPath := path.Join(execDir, WebConfigPath)
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("未找到配置文件，当前path:", configPath)
		os.Exit(1)
	}
}
