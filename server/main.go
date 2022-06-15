package server

import (
	"os"
	"os/signal"
	"sync-server/server/app"
	"sync-server/server/types"
	"syscall"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT,syscall.SIGTERM)
	app := app.NewApp(types.ConfigPath)

	err := app.RunServers()
	if err != nil {
		panic(err)
	}

	<- quit

	app.CloseServers()

}