package client

import (
	"context"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sourcegraph/jsonrpc2"
	"github.com/sovlookup/yunos/common"
)

type YunClient struct {
	ctx     context.Context
	h       *host.Host
	handler *YunClientHandler
}

type Yun interface {
	Event(msg any) error
	Close() error
}

type YunConn struct {
	ctx  context.Context
	conn *jsonrpc2.Conn
}

func (c *YunConn) Event(msg any) error {
	return c.conn.Notify(c.ctx, "event", msg)
}

func (c *YunConn) Close() error {
	return c.conn.Close()
}

func (server *YunClient) Connect(addr string) (Yun, error) {
	pi, err := peer.AddrInfoFromString(addr)

	if err != nil {
		return nil, err
	}

	(*server.h).Connect(server.ctx, *pi)

	if err != nil {
		return nil, err
	}

	stream, err := (*server.h).NewStream(server.ctx, pi.ID, common.PROTO_NAME)

	if err != nil {
		return nil, err
	}

	conn := &YunConn{
		server.ctx,
		jsonrpc2.NewConn(server.ctx, jsonrpc2.NewPlainObjectStream(stream), server.handler),
	}

	return conn, nil
}

func New(ctx context.Context, h *host.Host, meta *common.Meta, handler jsonrpc2.Handler) *YunClient {
	return &YunClient{ctx, h, &YunClientHandler{meta, handler}}
}
