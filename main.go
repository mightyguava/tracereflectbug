package main

import (
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	tracer := &httptrace.ClientTrace{
		ConnectStart: func(network, addr string) {
            log.Println("traced")
		},
	}
	req, err := http.NewRequest("GET", "http://localhost:9999/hello", nil)
	if err != nil {
		log.Printf("error creating request %v", err)
	}

	// Adding the same trace twice causes a stack overflow.
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), tracer))
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), tracer))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("request error: %v", err)
	}
	if res != nil && res.Body != nil {
		res.Body.Close()
	}
}
