package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1780115",
		Key:     "0db7222b8f2972a94ff8",
		Secret:  "5f25c249dfe6f6dccea8",
		Cluster: "ap2",
		Secure:  true,
	}

	app.Post("/api/messages", func(c fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)

		return c.JSON([]string{})
	})

	app.Listen(":8000")
}
