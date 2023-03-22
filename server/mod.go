package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"github.com/sourcegraph/jsonrpc2"
)

func Start(app *fiber.App, handler ...*Handler) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		println(c.RemoteAddr().String())
		ctx := context.Background()
		jrconn := jsonrpc2.NewConn(ctx, newWSObjectStream(c), newYunServerHandler(handler...))

		var meta any
		if err := jrconn.Call(ctx, "meta", nil, &meta); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		log.Println("meta", convertor.ToString(meta))
	}))

	log.Fatal(app.Listen(":3000"))
}
