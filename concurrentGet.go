package main

import (
	"log"
	"net/http"
	"time"
)

type links struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- links) {
	start := time.Now()

	if res, err := http.Get(url); err != nil {
		ch <- links{url, err, 0} //I got some error while fetching url
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- links{url, nil, t}
		res.Body.Close()
	}
}

func main() {
	resultsChan := make(chan links)

	urls := []string{
		"https://www.google.com",
		"https://www.amazon.in",
		"https://www.facebook.com",
	}
	for _, url := range urls {
		go get(url, resultsChan)
	}

	for range urls {
		r := <-resultsChan

		if r.err != nil {
			log.Printf(r.err)
		} else {
			log.Printf(r.url, r.latency)
		}
	}

}
