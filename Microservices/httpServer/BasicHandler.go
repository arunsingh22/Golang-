package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, r *http.Request) {

	//Routing from scratch
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(res, "Hello world from net/http Server :)")
		io.WriteString(res, "You are in the ROOT dir path")
	case "/about":
		io.WriteString(res, "You are in about section of the page")
	}
}

func main() {
	var d hotdog //d acts as a handler
	http.ListenAndServe(":8080", d)

}
