package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/WeasonTang/filetransfer/server"
)

//端口 port
var port = "27149"

func main() {
	go server.Run(port)
	cmd := startBrowser(port)
	chSignal := listenToInterrupt()
	//等待中断信号
	select {
	case <-chSignal:
		cmd.Process.Kill()
	}
}

//启动 Chrome
// 先写死路径，后面再照着 lorca 改
func startBrowser(port string) *exec.Cmd {
	chromePath := "C:/Program Files/Google/Chrome/Application/chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+port+"/static/index.html")
	cmd.Start()
	return cmd
}

//监听中断信号
func listenToInterrupt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
