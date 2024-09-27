package main

import (
	"github.com/thiagoaramizo/go-tarefas/config"
	"github.com/thiagoaramizo/go-tarefas/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// initialize config
	err := config.Init()
	if err != nil {
		logger.ErrorF("Error initializing config: %v", err)
		return
	}

	// initialize router
	router.Initialize()
}