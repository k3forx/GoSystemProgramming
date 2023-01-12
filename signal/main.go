package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	fmt.Println("Waiting SIGINT (CTRL+C)")
	<-signals
	fmt.Println("SIGINT arrived")
}

// func main() {
// 	ctx := context.Background()

// 	sigctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
// 	defer cancel()

// 	toctx, cancel2 := context.WithTimeout(ctx, 5*time.Second)
// 	defer cancel2()

// 	select {
// 	case <-sigctx.Done():
// 		fmt.Println("signal")
// 	case <-toctx.Done():
// 		fmt.Println("timeout")
// 	}
// }
