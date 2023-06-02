package poller

import (
	"os"
	"strconv"
)

const defaultIntervalInMs = 1000

type Config interface {
	IntervalInMs() int
}

type config struct {
	intervalInMs int
}

func NewConfig() Config {
	intervalInMs, err := strconv.Atoi(os.Getenv("POLL_INTERVAL_IN_MS"))
	if err != nil {
		intervalInMs = defaultIntervalInMs
	}
	return config{
		intervalInMs: intervalInMs,
	}
}

func (c config) IntervalInMs() int {
	return c.intervalInMs
}
