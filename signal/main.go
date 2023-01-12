package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()

	sigctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	toctx, cancel2 := context.WithTimeout(ctx, 5*time.Second)
	defer cancel2()

	select {
	case <-sigctx.Done():
		fmt.Println("signal")
	case <-toctx.Done():
		fmt.Println("timeout")
	}
}
