package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(handleAll)))
}

func handleAll(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}

	requestObj := &Request{
		Method:  req.Method,
		FullURL: req.URL.String(),
		Header:  Header(req.Header),
		Body:    string(body),
	}

	requestObj.Print()

	return
}
