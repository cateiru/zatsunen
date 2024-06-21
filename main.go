package main

import (
	"os"

	"github.com/cateiru/zatsunen/src/server"
)

// Set this variable at build time.
//
// Example:
//
//	go build  -ldflags="-X main.mode=prod"
var mode string = "local"

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	server.RunServer(mode, path)
}
