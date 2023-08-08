package main

import (
	"github.com/c12s/oort/internal/configs"
	"github.com/c12s/oort/internal/startup"
	"log"
)

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	err = startup.StartApp(config)
	if err != nil {
		log.Fatalln(err)
	}
}
