package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wgs .WaitGroup

// func network() {

// 	fmt.Println("doing some network stuff , chill")
// 	for i := 0; i < 10; i++ {

// 		time.Sleep(1 * time.Second)
// 		fmt.Println("network ..")
// 	}
// 	wgs.Done()

// }

// func ai() {

// 	fmt.Println("ai stuff , cool")
// 	for i := 0; i < 10; i++ {

// 		time.Sleep(1 * time.Second)
// 		fmt.Println("AI ..")
// 	}
// 	wgs.Done()
// }

// func testing() {

// 	fmt.Println("testing the code , oh very secure")
// 	for i := 0; i < 10; i++ {

// 		time.Sleep(1 * time.Second)
// 		fmt.Println("testing ..")
// 	}
// 	wgs.Done()

// }

// func main1() {

// 	wgs.Add(3)
// 	go network()
// 	go ai()
// 	go testing()

// 	wgs.Wait()
// }
