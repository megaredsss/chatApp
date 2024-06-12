package main

import (
	"fmt"
	"net/http"
)

// default page
func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
}
func main() {
	setupRoutes()
	fmt.Println("Hello world")
}
