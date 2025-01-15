package app

import (
	"time"
	"yordanluturyali/golang-auth-rest/exceptions"
	"github.com/gofiber/fiber/v2"
)

func NewRouter() *fiber.App {
	appRouter := fiber.New(fiber.Config{
		Prefork: true,
		AppName: "Golang Auth",
		IdleTimeout: 10 * time.Minute,
		ErrorHandler: exceptions.HandleError,
	})

	appRouter.Get("/", func(c * fiber.Ctx) error {
		return c.SendString("Halo")
	})

	return appRouter
}