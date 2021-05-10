package main

import "github.com/siampudan/learning/authsession/internal/server"

func main() {
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
