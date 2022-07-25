package jetstream

type Config interface {
}

type config struct {
}

//TODO: ovo je za prvu pomoc, izmeni posle
func NewDefaultJetStreamConfig() Config {
	return config{}
}
