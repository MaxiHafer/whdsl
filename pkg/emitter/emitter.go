package emitter

import (
	"time"

	"whdsl/pkg/ticker"
)

func NewEmitter(minInterval, maxInterval time.Duration, ) *Emitter {
	return &Emitter{
		ticker.NewRandomTicker(minInterval, maxInterval),
	}
}

type Emitter struct {
	ticker *ticker.RandomTicker
}
