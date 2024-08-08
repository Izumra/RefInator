package main

import (
	"github.com/Izumra/RefInator/app"
	configparser "github.com/Izumra/RefInator/utils/config_parser"
)

func main() {
	cfg := configparser.MustLoadByPath("config/settings.yaml")

	refinator := app.New(cfg)
	refinator.Refactor(cfg.FolderPath)
}
