package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Request struct {
	Method  string
	FullURL string
	Header  Header
	Body    string
}

func (r *Request) Print() {
	fmt.Println("URL:", r.FullURL)
	fmt.Println("Method:", r.Method)
	fmt.Println("Header:", r.Header.String())
	fmt.Println("Body:", r.Body)
}

type Header http.Header

func (h Header) String() (str string) {
	for key, vals := range h {
		str += (key + ": " + strings.Join(vals, ", "))
	}

	return
}
