package server

import (
	"github.com/VAISHAKH-GK/voting-tool-backend/internal/handlers"
	"github.com/VAISHAKH-GK/voting-tool-backend/pkg/sessions"
	"github.com/gofiber/fiber/v2"
)

func Run(port string) {
	var app = fiber.New()

	sessions.New()

	app.Post("/signup", handlers.Signup)
	app.Post("/otp", handlers.SubmitOTP)
	app.Post("/create-poll", handlers.CreatePoll)

	app.Listen(":" + port)
}
