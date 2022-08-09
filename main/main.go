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
	chChromeDie := make(chan struct{})
	chBackendDie := make(chan struct{})
	go server.Run(port)
	go startBrowser(port, chChromeDie, chBackendDie)
	chSignal := listenToInterrupt()
	//等待中断信号
	for {
		select {
		case <-chSignal:
			chBackendDie <- struct{}{}
		case <-chChromeDie:
			os.Exit(0)
		}
	}
}

//启动 Chrome
// 先写死路径，后面再照着 lorca 改
func startBrowser(port string, chChromeDie chan struct{},
	chBackendDie chan struct{}) {
	chromePath := "C:/Program Files/Google/Chrome/Application/chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+port+"/static/index.html")
	cmd.Start()
	go func() {
		<-chBackendDie
		cmd.Process.Kill()
	}()
	go func() {
		cmd.Wait()
		chChromeDie <- struct{}{}
	}()
}

//监听中断信号
func listenToInterrupt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
