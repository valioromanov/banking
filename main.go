package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	//log.Println("Starting application...")
	logger.Info("Starting application...")
	app.Start()
}
