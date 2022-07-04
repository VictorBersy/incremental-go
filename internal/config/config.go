package config

import (
	"time"
)

type Config struct {
	Tick time.Duration
}

func TickDuration() time.Duration {
	return 50 * time.Millisecond
}
