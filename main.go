package main

import "github.com/Tirivashe/hotel_management_backend/initializers"

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	println("Hello, world!")
}
