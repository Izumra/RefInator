package main

import (
	"log"

	"github.com/Izumra/RefInator/app/refinator"
	configparser "github.com/Izumra/RefInator/utils/config_parser"
)

func main() {
	cfg := configparser.MustLoadByPath("config/settings.yaml")

	refinator := refinator.New(cfg)
	err := refinator.MakeFolderCopy(cfg.FolderPath)
	if err != nil {
		log.Println(err)
	}

	refinator.Refactor(cfg.FolderPath + "_copy")
}

//func main() {
//	function, err := swift.GenFunction()
//	if err != nil {
//		log.Println(err)
//	}
//
//	log.Printf("Сгенерированная функция: \n%v", function)
//}
//
