package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

type Person struct {
	Name string
	Age int
}

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			fmt.Printf("Done err is: %v\n", ctx.Err())
			return ctx.Err()
		}
	}
}

func main() {
	//rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10; i ++ {
		fmt.Println(rand.Int())
	}
}
