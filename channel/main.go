package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub1()")
	go sub1()
	time.Sleep(2 * time.Second)
}

func sub1() {
	fmt.Println("sub1() is running")
	time.Sleep(time.Second)
	fmt.Println("sub1() is finished")
}
