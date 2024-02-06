package main

import (
	"fmt"

	"github.com/trickaugusto/gopportunities/config"
	"github.com/trickaugusto/gopportunities/router"
)

func main() {
	// inicialize configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Call the router package Initialize function
	router.Initialize()
}
