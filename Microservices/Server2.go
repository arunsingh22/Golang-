package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to First Server in Go!")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ooops", http.StatusBadRequest)
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Oops"))
			return
		}
		fmt.Fprintf(w, "Hello %s", b)
	}

	http.HandleFunc("/", f)

	//Creating a HTTP server from ground up.
	s := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      f, //fix this.
	}
	s.ListenAndServe()

	//Gracefully shutdowns my server after 20 seconds when my handlers are all free.
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	s.Shutdown(ctx)

}
