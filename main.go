package main

import (
	"fmt"
	"time"
)

func main() {
	go run()
	time.Sleep(5 * time.Second)
}

func run() {
	defer recoverPanic()
	for i := 0; i < 5; i++ {
		fmt.Println("going on")
		if i == 2 {
			panic("panicked")
		}
	}
}

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic")
	}
}
