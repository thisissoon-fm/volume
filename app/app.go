package app

import (
	"os"
	"os/signal"
	"syscall"

	"volume/ipc/websocket"
	"volume/log"
	"volume/volume"
)

type App struct {
}

func (a *App) Run() error {
	log.Info("app start")
	defer log.Info("app exit")
	// Load Volume
	volume.LoadState()
	defer volume.SaveSate()
	// Start websocket client
	ws := websocket.New(websocket.Config{})
	go ws.Connect()
	defer ws.Close()
	// Wait OS signal to exit
	sigC := make(chan os.Signal, 1)
	signal.Notify(
		sigC,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	signal := <-sigC
	log.WithField("signal", signal).Debug("received os signal")
	return nil
}

func New() *App {
	return &App{}
}
