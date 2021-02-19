package main

import (
	"fmt"
	"net/http"
	"github.com/ravindrags24/sto-repo/go-project/pkg"
)

const portNumber = ":8080"

// main is the main application function.
func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
