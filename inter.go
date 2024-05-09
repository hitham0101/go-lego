package main

// import (
// 	"fmt"
// 	"time"
// )

// type stub interface {
// 	sharp()
// }

// type sword struct {
// 	name string
// }
// type knife struct {
// 	name string
// }

// func (s sword) sharp() {

// 	fmt.Println("swooooord power 100%")
// }
// func (s knife) sharp() {

// 	fmt.Println("knife power 20%")
// }

// func help(attack stub) {

// 	if sword, ok := attack.(sword); ok {
// 		fmt.Println("here is " + sword.name + " How can I help you sir")
// 	}

// 	if knife, ok := attack.(knife); ok {
// 		fmt.Println("here is " + knife.name + " How can I help you sir")
// 	}

// 	attack.sharp()
// }

// func main() {

// 	weapon1 := sword{"spiderman"}
// 	help(weapon1)

// 	for i := 0; i < 10; i++ {
// 		var char = "="
// 		char = char + "="
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Print(char)
// 	}
// 	fmt.Println("  ")

// 	weapon2 := knife{"jocker"}
// 	help(weapon2)

// }
