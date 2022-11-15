package main

import (
	"net/http"

	"github.com/tquigly1287/go-dev-workspace/mortgage"
)

func main() {
	http.HandleFunc("/", mortgage.Greet)
	http.ListenAndServe(":8080", nil)
}