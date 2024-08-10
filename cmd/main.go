package main

import (
	"log"

	"github.com/Izumra/RefInator/app"
	configparser "github.com/Izumra/RefInator/utils/config_parser"
)

func main() {
	cfg := configparser.MustLoadByPath("config/settings.yaml")

	refinator := app.New(cfg)
	err := refinator.MakeFolderCopy(cfg.FolderPath)
	if err != nil {
		log.Println(err)
	}

	refinator.Refactor(cfg.FolderPath + "_copy")
}
