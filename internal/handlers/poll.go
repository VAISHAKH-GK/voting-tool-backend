package handlers

import (
	"github.com/VAISHAKH-GK/voting-tool-backend/pkg/poll"
	"github.com/VAISHAKH-GK/voting-tool-backend/pkg/sessions"
	"github.com/gofiber/fiber/v2"
)

func CreatePoll(c *fiber.Ctx) error {
	var username string
	if username = sessions.Get(c, "username").(string); username == "" {
		return c.Status(401).JSON("Not Logged In")
	}

	var poll poll.Poll
	c.BodyParser(&poll)
	poll.AddPoll()
	return c.JSON("Poll added")
}
