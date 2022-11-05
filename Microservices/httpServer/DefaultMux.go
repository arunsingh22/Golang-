package main

import (
	"io"
	"net/http"
)

// type hotdog int
// type hotcat int

// func (d hotdog) ServeHTTP(res http.ResponseWriter, r *http.Request) {
// 	io.WriteString(res, "Hello from DOG")
// }

// func (c hotcat) ServeHTTP(res http.ResponseWriter, r *http.Request) {
// 	io.WriteString(res, "Hello from CAT")
// }

func d(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Hello from DOG")
}

func c(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Hello from CAT")
}

func main() {
	// var d hotdog //d acts as a handler
	// var c hotcat

	// mux := http.NewServeMux()
	// mux.Handle("/dog/", d)
	// mux.Handle("/cat", c)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)

}
