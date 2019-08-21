package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var repeat = func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valuedStream := make(chan interface{})
	go func() {
		defer close(valuedStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valuedStream <- v:
				}
			}
		}
	}()
	return valuedStream
}

var repeatFn = func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for {
			select {
			case <-done:
				return
			case outStream <- fn():
			}
		}
	}()
	return outStream
}

var take = func(done <-chan interface{}, valuedStream <-chan interface{}, num int) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case outStream <- <-valuedStream:
			}
		}
	}()
	return outStream
}

var toString = func(done <-chan interface{}, valuedStream <-chan interface{}) <-chan string {
	outStream := make(chan string)
	go func() {
		defer close(outStream)
		for v := range valuedStream {
			select {
			case <-done:
				return
			case outStream <- v.(string):
			}
		}
	}()
	return outStream
}

var fanIn = func(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	muliplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go muliplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

var orDone = func(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

var bridge = func(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if ok == false {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}
			for val := range orDone(done, stream) {
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func example2() {
	done := make(chan interface{})
	defer close(done)
	for num := range take(done, repeat(done, 1, 2, 3), 10) {
		fmt.Printf("%v", num)
	}
	rand := func() interface{} {
		rand.Seed(time.Now().UnixNano())
		return rand.Int()
	}

	log.Println()
	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}

	log.Println()
	var message string
	for num := range toString(done, take(done, repeat(done, "I", "am."), 3)) {
		message += num
	}
	fmt.Println(message)
}
