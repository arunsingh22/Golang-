package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to First Server in Go!")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ooops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s", b)

	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, _ *http.Request) {
		log.Println("Good Bye Folks from Go!")
	})

	http.ListenAndServe(":8000", nil)
}
