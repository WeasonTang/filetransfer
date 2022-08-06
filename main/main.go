package main

import (
	"os"
	"os/signal"

	"github.com/zserge/lorca"
)

func main() {
	var ui lorca.UI
	ui, _ = lorca.New("https://baidu.com", "", 800, 600,
		"--disable-sync", "--disable-translate")
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal)
	select {
	case <-ui.Done():
	case <-chSignal:
	}
	ui.Close()
}
