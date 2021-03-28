package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeout int = 5000

type Response struct {
	Err  error
	Resp string
}

func barrier(urls ...string) {
	numOfRequests := len(urls)
	in := make(chan Response, numOfRequests)
	defer close(in)

	responses := make([]Response, numOfRequests)

	for _, uri := range urls {
		go makeRequest(
			in,
			uri,
		)
	}

	var hasError bool
	for i := 0; i < numOfRequests; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

func makeRequest(out chan<- Response, url string) {
	res := Response{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeout) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	res.Resp = string(byt)
	out <- res
}
