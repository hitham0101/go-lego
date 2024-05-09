package main

// import (
// 	"fmt"
// 	"sync"
// )

// var weg sync.WaitGroup

// func sen(ch chan int) {

// 	ch <- 2
// 	ch <- 3
// 	weg.Done()

// }
// func recv(ch chan int) {

// 	val := <-ch
// 	fmt.Println(val)
// 	// val = <-ch
// 	// fmt.Println(val)

// 	weg.Done()
// }

// func chee() {

// 	weg.Add(2)
// 	fmt.Println("HEY")
// 	ch := make(chan int, 1)

// 	go sen(ch)
// 	go recv(ch)

// 	weg.Wait()
// }
