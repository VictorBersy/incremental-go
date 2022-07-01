package config

import (
	"time"
)

type Config struct {
	Tick time.Duration
}

func Load() (config Config) {
	return Config{
		Tick: 1 * time.Millisecond,
	}
}
