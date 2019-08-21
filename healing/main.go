package main

import (
	"time"
)

type startGoroutineFn func(done <-chan interface{}, pulseInterval time.Duration) (heartbeat <-chan interface{})

var newSteward = func(timeout time.Duration, startGoroutine startGoroutineFn) startGoroutineFn {
	return func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
		heartbeat := make(chan interface{})
		go func() {
			defer close(heartbeat)

			var wardDone chan interface{}
			var wardHeartbeat <-chan interface{}
			startWard := func() {
				wardDone = make(chan interface{})
				wardHeartbeat = startGoroutine(or)
			}
		}()
	}
}

func main() {

}
