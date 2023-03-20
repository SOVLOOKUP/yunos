package server

import (
	"context"
	"log"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/sourcegraph/jsonrpc2"
)

type yunServerHandler struct{}

func (h *yunServerHandler) Handle(ctx context.Context, c *jsonrpc2.Conn, r *jsonrpc2.Request) {
	switch r.Method {
	case "event":
		params, _ := r.Params.MarshalJSON()
		// todo receive event
		log.Println(convertor.ToString(params))
		return
	default:
		err := &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: "Method not found"}
		if err := c.ReplyWithError(ctx, r.ID, err); err != nil {
			log.Println(err)
			return
		}
	}
}
