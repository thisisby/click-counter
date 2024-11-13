package main

import (
	"click-counter/internal/app"
	"click-counter/internal/config"
	"click-counter/pkg/logger"
)

func init() {
	logger.InitZeroLogger()
	config.Config.MustInitializeConfig()
}

func main() {
	app.MustRun()
}
