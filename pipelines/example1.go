package main

import "fmt"

func example1() {
	generator := func(done chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(done chan interface{}, input <-chan int, multiplier int) <-chan int {
		outStream := make(chan int)
		go func() {
			defer close(outStream)
			for i := range input {
				select {
				case <-done:
					return
				case outStream <- i * multiplier:
				}
			}
		}()
		return outStream
	}

	add := func(done chan interface{}, input <-chan int, additive int) <-chan int {
		outStream := make(chan int)
		go func() {
			defer close(outStream)
			for i := range input {
				select {
				case <-done:
					return
				case outStream <- (i + additive):
				}
			}
		}()
		return outStream
	}

	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, intStream, 1), 2)

	for v := range pipeline {
		if v == 6 {
		}
		fmt.Println(v)
	}
}
