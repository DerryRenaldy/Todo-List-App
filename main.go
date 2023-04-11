package main

import (
	"github.com/DerryRenaldy/Todo-List-App/configs"
	"github.com/DerryRenaldy/Todo-List-App/server"
	"github.com/DerryRenaldy/logger/logger"
)

func main() {
	cfg := configs.Cfg
	log := logger.New(cfg.App.Environment, cfg.App.AppName, cfg.App.LogLevel)
	s := server.NewService(cfg, log)
	s.Start()
}
