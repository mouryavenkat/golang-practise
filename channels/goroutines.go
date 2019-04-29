package main

import (
	"fmt"
	"time"
)

func helloGoroutine(name string, ch chan string) {
	fmt.Println("Hello ", name)
	time.Sleep(2 * time.Second)
	ch <- name
}

// Asynchronous way (Use gorutines)

func main() {
	ch := make(chan string)
	go helloGoroutine("mourya", ch)
	go helloGoroutine("vasavi", ch)
	fmt.Println("Exiting the process")
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
	// So we told that until all are received we dont do any action. Like await Promise.all(Promises) in node

	fmt.Println("Goroutines completed execution")
	// By default doesnt wait for the gorutines to complete. Here comes channels
	// One more thing to remember is that , you cant control the execution of the goroutine, Its handled by Go internally.
}

// Synchronous way

// func main() {
// 	helloGoroutine("mourya")
// 	helloGoroutine("vasavi")
// 	fmt.Println("Exiting the process")
// }
