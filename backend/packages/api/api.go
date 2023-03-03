package api

import (
	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
)

var server *fiber.App

func StartServer() {
	log.Info("Starting server")
	server = httpServer()
	serverErr := server.Listen(":3003")

	if serverErr != nil {
		log.WithField("reason", serverErr.Error()).Fatal("Server error")
	}
}

func StopServer() {
	if server != nil {
		log.Info("Stopping server")
		err := server.Shutdown()
		if err != nil {
			log.WithField("reason", err.Error()).Fatal("Shutdown server error")
		}
	}
}
