package main

import (
	"os"

	"github.com/sodefrin/appengine-boilerplate/app"
	"github.com/sodefrin/appengine-boilerplate/log"
	"go.uber.org/zap"
)

var (
	exitOK    = 0
	exitError = 1
)

func main() {
	log.Logger.Info("started")

	err := app.Run()
	if err != nil {
		log.Logger.Error("failed", zap.Error(err))
		os.Exit(exitError)
	}

	log.Logger.Info("stopped")
	os.Exit(exitOK)
}
