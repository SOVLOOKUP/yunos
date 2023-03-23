package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sovlookup/yunos/server"
)

func main() {
	app := fiber.New()
	server.Start(app)
}
