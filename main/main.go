package main

import (
	_ "github.com/treeder/go-polymer"
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:3000", nil)
}