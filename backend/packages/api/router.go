package api

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	controllers "github.com/comfucios/puppet-installer/backend/packages/api/controllers"
)

//go:embed ui/*
var staticFolder embed.FS

func httpServer() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin",
		AllowMethods:     "POST, OPTIONS, GET, PUT",
		ExposeHeaders:    "Set-Cookie",
	}))

	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(staticFolder),
		PathPrefix: "ui",
	}))

	api := app.Group("/api")

	api.Get("/ping", controllers.Pong)

	infra := api.Group("/infra")
	infra.Post("/install", controllers.Infra)

	return app
}
