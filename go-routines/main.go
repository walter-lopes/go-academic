package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello"

	wg.Add(1)

	// generates copy wh en passed args in anonymous functions
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)

	wg.Wait()

	//fmt.Println("Aloha")
	msg = "Goodbye"
	//time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
