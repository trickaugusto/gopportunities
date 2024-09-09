package main

import (
	"github.com/trickaugusto/gopportunities/config"
	"github.com/trickaugusto/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// inicialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	// Call the router package Initialize function
	router.Initialize()
}
