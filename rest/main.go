package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "weclome man")

}

func create(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "ParseForm() error", http.StatusBadRequest)
		return
	}

	// Get the file name from the form data
	fileName := r.FormValue("file_name")

	// Create a new logger
	logger := log.New(os.Stdout, "PREFIX: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Creating file:", fileName)

	// Create the file with the provided file name
	_, err = os.Create(fileName)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}

	// Write a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File created successfully"))
}

func main() {
	fmt.Println("start")
	http.HandleFunc("/home", home)
	http.HandleFunc("/create", create)
	http.ListenAndServe(":8080", nil)
}
