package main

import (
	"log"

	"github.com/Tirivashe/hotel_management_backend/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	server := NewApiServer(":8000")
	if err := server.Start(); err != nil {
		log.Fatal("Cannot start server")
	}
	println("Hello, world!")
}
