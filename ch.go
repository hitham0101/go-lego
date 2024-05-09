package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup // 1

// func send(ch chan int) {
// 	defer wg.Done() // 3
// 	ch <- 40
// 	ch <- 20
// 	fmt.Println("send finished")
	
// 	close(ch) // Closing the channel after sending all values

// }

// func rec(ch chan int) {
// 	defer wg.Done() // 3
// 	val, ok := <-ch // Receiving data from the channel
// 	fmt.Println(ok)
// 	fmt.Println("reccive finished", val)
// 	if !ok {
// 		fmt.Println("Channel closed, exiting")
// 		return
// 	}
// 	val, ok = <-ch // Receiving data from the channel
// 	fmt.Println("reccive finished", val)
// 	if !ok {
// 		fmt.Println("Channel closed, exiting")
// 		return
// 	}

// }
// func ch() {

// 	ch := make(chan int)

// 	wg.Add(2)   // 2
// 	go send(ch) // *
// 	go rec(ch)  // *

// 	wg.Wait() // 4
// 	fmt.Println("main finished")
// }
