package main

import (
	"fmt"
	"net/http"
)

// (2)How to make URL root path to this new function. / Mapping
// Then bind our server to a network port, so it can start taking requests.

// do the mapping in the main function so that it takes place as soon as we run the program
func main() {
	fmt.Println("main")
	// http.HandleFunc -->  routes url paths to function handlers
	// first arg is path
	// second is the function handler
	http.HandleFunc("/", helloHandler)

	// (3)To finish our server we will use the ListenAndServer function
	// starts the server
	// and listens/binds on the given port
	// second arg is the multiplexer --> used to dispatch the http request to the handler
	// // can just use nil --> sets a default value
	// // // usually won't need to set this value ever
	http.ListenAndServe(":8080", nil)
}

// (1) The Function Handler
// the request handler function
// must take 2 arguments
// http.ResponseWrite --> used for sending data back in the response
// *http.Request --> includes information about the request, including data sent by the client
// // * indicated pointer to a value
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// takes 2 arguments
	// first is the response writer
	// second is the text we want to send
	fmt.Fprintf(w, "Hello world\n")
}

// (4) when we run the program, the server starts up, and listens on port 8080, and waits for incoming connections
// go run main.go
// see localhost:8080
