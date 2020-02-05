package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; 1++ {
		wg.Add(2)
		go sayHello2()
		go increment()
	}

	wg.Wait()
}

func sayHello2() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
