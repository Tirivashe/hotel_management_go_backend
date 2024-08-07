package main

import (
	"github.com/Tirivashe/hotel_management_backend/routes"
	"github.com/gofiber/fiber/v2"
)

type ApiServer struct {
	mux  *fiber.App
	port string
}

func NewApiServer(port string) *ApiServer {
	mux := fiber.New()
	routes.LoadRoutes(mux)
	return &ApiServer{
		mux:  mux,
		port: port,
	}
}

func (s *ApiServer) Start() error {
	return s.mux.Listen(s.port)
}
