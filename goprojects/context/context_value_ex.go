package main

import (
	"context"
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func ContextValueEx() {

	wg2.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9)
	go square(ctx)

	wg2.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
	wg2.Done()
}
