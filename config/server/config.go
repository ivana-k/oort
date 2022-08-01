package server

import "os"

type Config interface {
	Port() string
}

type config struct {
	port string
}

func NewConfig() Config {
	return config{
		port: os.Getenv("OORT_PORT"),
	}
}

func (c config) Port() string {
	return c.port
}
