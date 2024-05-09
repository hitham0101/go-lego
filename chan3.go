package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wek sync.WaitGroup

// func senv(che chan<- int, v int) {

// 	che <- v
// 	wek.Done()
// }

// func rev(che <-chan int) {

// 	val := <-che
// 	fmt.Println(val)
// 	wek.Done()

// }
// func cheeeee() {
// 	wek.Add(2)
// 	ch := make(chan int)
// 	go senv(ch, 2)
// 	go rev(ch)
// 	wek.Wait()

// }
