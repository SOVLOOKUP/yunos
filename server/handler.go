package server

import (
	"context"
	"log"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/sourcegraph/jsonrpc2"
)

type Handler interface {
	Handle(context.Context, *jsonrpc2.Conn, *jsonrpc2.Request) error
}

type yunServerHandler struct {
	customHandler []*Handler
}

func (h *yunServerHandler) Handle(ctx context.Context, c *jsonrpc2.Conn, r *jsonrpc2.Request) {
	switch r.Method {
	case "event":
		params, _ := r.Params.MarshalJSON()
		// todo receive event
		log.Println(convertor.ToString(params))
		return
	default:
		for _, v := range h.customHandler {
			if (*v).Handle(ctx, c, r) != nil {
				continue
			} else {
				return
			}
		}
		err := &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: "Method not found"}
		if err := c.ReplyWithError(ctx, r.ID, err); err != nil {
			log.Println(err)
			return
		}
	}
}

func newYunServerHandler(customHandler ...*Handler) *yunServerHandler {
	return &yunServerHandler{
		customHandler,
	}
}
