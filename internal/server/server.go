package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(),
	}

	server.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Replace with your frontend domain
		AllowMethods: "POST",
		AllowHeaders: "*",
	}))

	return server
}
