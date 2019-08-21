package main

import (
	"log"
	"sync"
)

func main() {
	result := make(chan int)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go doWork(result, &wg)
	}
	result <- 1
	wg.Wait()
}

func doWork(result <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case v := <-result:
		log.Println(v)
	}
}
