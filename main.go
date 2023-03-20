package main

import (
	"context"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/sovlookup/yunos/client"
	"github.com/sovlookup/yunos/server"
)

func main() {
	// server

	ctx := context.Background()
	s, err := libp2p.New()
	if err != nil {
		panic(err)
	}
	yuns := server.New(ctx, &s)
	// println(yuns.Addr())

	go func() {
		for i := 0; i < 1000; i++ {
			err := newClient(yuns.Addr())
			if err != nil {
				println(err)
			}
			// println(i)
		}
	}()

	select {}
}

func newClient(addr string) error {
	// client
	ctx := context.Background()
	c, err := libp2p.New()
	if err != nil {
		return err
	}
	yunc := client.New(ctx, &c, nil, &myHandler{})
	conn, err := yunc.Connect(addr)
	if err != nil {
		return err
	}

	if conn.Event("some msg") != nil {
		return err
	}
	return nil
}

type myHandler struct{}

// Handle implements the jsonrpc2.Handler interface.
func (h *myHandler) Handle(ctx context.Context, c *jsonrpc2.Conn, r *jsonrpc2.Request) {
	switch r.Method {
	case "sayHello":
		if err := c.Reply(ctx, r.ID, "hello world"); err != nil {
			log.Println(err)
			return
		}
	default:
		err := &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: "Method not found"}
		if err := c.ReplyWithError(ctx, r.ID, err); err != nil {
			log.Println(err)
			return
		}
	}
}
