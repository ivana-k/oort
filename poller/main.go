package main

import (
	"github.com/c12s/oort/poller/app"
	"github.com/c12s/oort/poller/config"
)

func main() {
	conf := config.NewConfig()
	app.Run(conf)
}
