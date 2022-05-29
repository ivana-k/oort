package main

import (
	"github.com/c12s/oort/app"
	"github.com/c12s/oort/config"
)

func main() {
	config := config.NewDefaultConfig()
	app.Run(config)
}
