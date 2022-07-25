package server

type Config interface {
	Host() string
	Port() string
}

type config struct {
	host string
	port string
}

func NewDefaultServerConfig() Config {
	return config{
		host: "",
		port: "8000",
	}
}

func (c config) Host() string {
	return c.host
}

func (c config) Port() string {
	return c.port
}
