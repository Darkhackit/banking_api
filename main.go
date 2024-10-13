package main

import (
	"github/Darkhackit/banking_api/app"
	"github/Darkhackit/banking_api/logger"
)

func main() {
	//log.Println("Starting our banking api")
	logger.Info("Starting Banking API")
	app.Start()
}
