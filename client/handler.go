package client

import (
	"context"
	"log"

	"github.com/sourcegraph/jsonrpc2"
	"github.com/sovlookup/yunos/common"
)

type YunClientHandler struct {
	meta    *common.Meta
	handler jsonrpc2.Handler
}

func (h *YunClientHandler) Handle(ctx context.Context, c *jsonrpc2.Conn, r *jsonrpc2.Request) {
	switch r.Method {
	case "meta":
		if err := c.Reply(ctx, r.ID, h.meta); err != nil {
			log.Println(err)
			return
		}
	default:
		h.handler.Handle(ctx, c, r)
	}
}
