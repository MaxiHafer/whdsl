package ticker

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func NewRandomTicker(min, max time.Duration) *RandomTicker {
	ticker := &RandomTicker{
		C: make(chan time.Time),
		stopc: make(chan struct{}),
		min: min.Nanoseconds(),
		max: max.Nanoseconds(),
	}

	go ticker.loop()
	return ticker
}

type RandomTicker struct {
	C chan time.Time
	stopc chan struct{}
	min int64
	max int64
}

func (rt *RandomTicker) loop() {
	ticker := time.NewTicker(rt.nextInterval())
	for {
		select {
		case <-rt.stopc:
			ticker.Stop()
			return
		case <-rt.C:
			select {
			case rt.C <- time.Now():
				ticker.Stop()
				ticker = time.NewTicker(rt.nextInterval())
			default:
			}
		}
	}
}

func (rt *RandomTicker) nextInterval() time.Duration {
	interval := rand.Int63n(rt.max - rt.min) + rt.min

	return time.Duration(interval) * time.Second
}

func (rt *RandomTicker) Stop() {
	close(rt.stopc)
	close(rt.C)
}
