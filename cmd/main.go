package main

import (
	"log"
	configPkg "mobidevtestProject/cmd/config"
	"mobidevtestProject/logs"
	"mobidevtestProject/pkg"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config := configPkg.CreateConfig()
	if err := configPkg.ReadConfig("cmd/config/devConfig.json", config); err != nil {
		log.Fatal(err)
	}
	logFile := logs.CreateLogFile()
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(logFile)
	logger := logs.NewLogger(logFile)
	server := pkg.InitServer(config, logger)
	err := server.StartServer()
	if err != nil {
		log.Fatal(err)
	}

}
