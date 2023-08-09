package services

type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, tags []string) error
	Invalidate(tags []string) error
	Stop()
}
