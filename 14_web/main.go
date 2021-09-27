package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Hello World</h1>")
}

func about(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>About</h1>")
}

func main() {
	// routing
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	// serve at http://localhost:3000
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
	// ctrl + C to stop serving
}
