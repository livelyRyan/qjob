package main

import (
	"github.com/spf13/pflag"
	"qjob/master"
	"runtime"
	"time"
)

var (
	configPath string
)

func main() {
	initEnv()
	initParams()

	err := master.InitGlobalConfig(configPath)
	if err != nil {
		panic(err)
	}

	err = master.InitJobMgr()
	if err != nil {
		panic(err)
	}

	// 初始化 apiServer
	err = master.InitApiServer()
	if err != nil {
		panic(err)
	}

	// 让主线程不退出
	for true {
		time.Sleep(time.Second)
	}
}

func initEnv() {
	// 初始化 process 数量
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 解析命令行参数
func initParams() {
	pflag.StringVarP(&configPath, "configPath", "f", "./master.json", "Master config file path")
	pflag.Parse()
}
