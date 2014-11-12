// +build !appengine
// Above is a special build command: https://blog.golang.org/the-app-engine-sdk-and-workspaces-gopath

package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:8080", nil)
}
