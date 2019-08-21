package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("Cannot print greeting: %v\n", err)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("Cannot print farewell: %v", err)
			return
		}
	}()
	wg.Wait()
}

func printGreeting(ctx context.Context) error {
	greetign, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greetign)
	return nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewel(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}

	return "", fmt.Errorf("unsupported locale")
}

func genFarewel(ctx context.Context) (string, error) {
	switch local, err := locale(ctx); {
	case err != nil:
		return "", err
	case local == "EN/US":
		return "goodbye", nil
	}

	return "", fmt.Errorf("unsuppored locale")
}

func locale(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(10 * time.Second):
	}
	return "EN/US", nil
}
