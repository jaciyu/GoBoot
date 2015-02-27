package main

import (
	"GoBoot/src/controllers"
	"GoBoot/src/lib"
	"flag"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

func main() {
	defer func() {
		managePid(false) //删除pid文件
		println("Server Exit")
	}()
	envInit()
	managePid(true) //生成pid文件
	startServer()
}

//所有初始化操作
func envInit() {
	// println(os.Args[0])
	os.Chdir(path.Dir(os.Args[0]))
	confiPath := flag.String("f", "../conf/app.conf", "config file")
	flag.Parse()
	if *confiPath == "" {
		panic("config file missing")
	}
	lib.Conf.Init(*confiPath)
	lib.Logger.Init(lib.Conf.Get("log_root"), lib.Conf.GetInt("log_level"))
	//如果需要访问mongo参考README中相关说明，并打开以下注释
	//lib.MongoSrv.Init(lib.Conf.Get("mongodb_host"), lib.Conf.GetInt("mongodb_timeout"), lib.Conf.GetInt("mongodb_mode"))
}

//启动httpserver
func startServer() {
	server := lib.NewHttpServer("", lib.Conf.GetInt("http_port"),
		lib.Conf.GetInt("http_timeout"),
		lib.Conf.GetBool("pprof_enable"))
	//每个 controller需要在这里注册, HttpApiServer能够根据请求参数自动路由
	server.AddController(&controllers.MonitorController{})
	server.AddController(&controllers.DemoController{})
	println("Server Start")
	server.Run()
}

//生成/删除当前进程id文件
func managePid(create bool) {
	pidFile := lib.Conf.Get("app_pid_file")
	if create {
		pid := os.Getpid()
		pidString := strconv.Itoa(pid)
		ioutil.WriteFile(pidFile, []byte(pidString), 0777)
	} else {
		os.Remove(pidFile)
	}
}
