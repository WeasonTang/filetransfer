package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/WeasonTang/filetransfer/server"
)

func main() {

	//端口 port
	port := "27149"
	//启动 gin 服务
	go server.Run(port)

	//启动 Chrome
	// 先写死路径，后面再照着 lorca 改
	chromePath := "C:/Program Files/Google/Chrome/Application/chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+port+"/static/index.html")
	cmd.Start()

	//监听中断信号
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	//等待中断信号
	select {
	case <-chSignal:
		cmd.Process.Kill()

	}
}
