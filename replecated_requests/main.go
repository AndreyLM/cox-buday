package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var doWork = func(ctx context.Context, id int, wg *sync.WaitGroup, result chan<- int) {
	started := time.Now()
	defer wg.Done()
	simulatedLoadTime := time.Duration(1*rand.Intn(5)) * time.Second

	select {
	case <-ctx.Done():
	case <-time.After(simulatedLoadTime):
	}

	select {
	case <-ctx.Done():
	case result <- id:
	}

	took := time.Since(started)
	if took < simulatedLoadTime {
		took = simulatedLoadTime
	}
	fmt.Printf("%v took %v\n", id, took)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan int)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go doWork(ctx, i, &wg, result)
	}
	firstReturned := <-result
	cancel()
	wg.Wait()
	fmt.Printf("Received an answer from #%v\n", firstReturned)
}
