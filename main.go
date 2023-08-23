package main

import (
	"goscrape/server"
)

func main() {
	serverErr := server.InitializeServer()
	if serverErr != nil {
		panic(serverErr)
	}
}
