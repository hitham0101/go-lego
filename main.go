package main

import (
	"log"
	"os"
)

func main() {
	// file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	logger := log.New(os.Stdout, "PREFIX: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("This is a custom log message")
}
