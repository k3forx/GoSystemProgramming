package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("------------------------")
	fmt.Println("start sub1()")
	go sub1()
	time.Sleep(2 * time.Second)

	fmt.Println("------------------------")

	pn := primeNumber()
	for n := range pn {
		fmt.Println(n)
	}

	fmt.Println("------------------------")
	idCh := make(chan int)
	ids := []int{1, 2, 3, 4, 5}
	go func() {
		// 1秒ごとにチャネルにidを送信
		// すべて送信したら閉じる
		defer close(idCh)
		for _, id := range ids {
			time.Sleep(1 * time.Second)
			idCh <- id
		}
	}()
loop:
	for {
		select {
		case id, ok := <-idCh:
			// チャネルが閉じられたら終了
			if !ok {
				break loop
			}
			fmt.Println(id)
		default:
			fmt.Println("no value...")
		}
	}
	fmt.Println("Do something...")
	fmt.Println("Done!!")
}

func sub1() {
	fmt.Println("sub1() is running")
	time.Sleep(time.Second)
	fmt.Println("sub1() is finished")
}

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 100; i += 2 {
			l := int(math.Sqrt(float64(i)))
			var found bool
			for j := 3; j < l+1; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}
