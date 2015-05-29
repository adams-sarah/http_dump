package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Method  string
	FullURL string
	Header  http.Header
	Body    []byte
}

func main() {
	mux := newMux()
	log.Fatal(http.ListenAndServe(":5000", mux))
}

func handleAll(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	requestObj := &Request{
		Method:  req.Method,
		FullURL: req.URL.String(),
		Header:  req.Header,
		Body:    body,
	}

	return
}
