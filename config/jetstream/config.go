package kafka

type Config interface {
}

type config struct {
}

//TODO: ovo je za prvu pomoc, izmeni posle
func NewKafkaConfig() Config {
	return config{}
}
