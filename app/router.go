package app

import (
	"time"
	"yordanluturyali/golang-auth-rest/controllers"
	"yordanluturyali/golang-auth-rest/exceptions"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(authController controllers.AuthController) *fiber.App {
	appRouter := fiber.New(fiber.Config{
		Prefork: true,
		AppName: "Golang Auth",
		IdleTimeout: 10 * time.Minute,
		ErrorHandler: exceptions.HandleError,
	})

	return appRouter
}