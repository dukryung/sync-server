package app

import (
	"fmt"
	"sync-server/server/sync"
	"sync-server/server/types"
)

type App struct {
	servers []types.Server
	config  types.AppConfig
}

func NewApp(configPath string) *App {

	app := App{}
	err := app.config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	app.initServers()

	return &app
}

func (app *App) initServers() {
	if app.config.Sync.Enable {
		app.servers = append(app.servers, sync.NewServer(app.config.Sync))
	}
}

func (app *App) RunServers() error {
	if len(app.servers) == 0 {
		return fmt.Errorf("empty servers")
	}

	for _, server := range app.servers {
		server.Run()
	}

	return nil
}

func (app *App) CloseServers() {
	for _, server := range app.servers {
		server.Close()
	}
}
