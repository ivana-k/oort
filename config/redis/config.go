package redis

type Config interface {

}

func NewRedisConfig() Config {
	return config{}
}

type config struct {

}
