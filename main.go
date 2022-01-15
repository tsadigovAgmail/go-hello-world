package main

import (
	"net/http"

	"github.com/l0rd/go-hello-world/helloworld"
)

func main() {
	http.HandleFunc("/", helloworld.Greet)
	http.ListenAndServe(":8080", nil)
}