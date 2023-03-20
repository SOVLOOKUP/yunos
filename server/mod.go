package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/sovlookup/yunos/common"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/multiformats/go-multiaddr"
	"github.com/sourcegraph/jsonrpc2"
)

type YunServer struct {
	ctx  context.Context
	h    *host.Host
	conn map[string]*jsonrpc2.Conn
	meta map[string]*common.Meta
}

func (server *YunServer) initS() {
	(*server.h).SetStreamHandler(common.PROTO_NAME, func(stream network.Stream) {
		jrconn := jsonrpc2.NewConn(server.ctx, jsonrpc2.NewPlainObjectStream(stream), &yunServerHandler{})
		server.conn[stream.ID()] = jrconn

		var meta common.Meta
		if err := jrconn.Call(server.ctx, "meta", nil, &meta); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		log.Println("meta", convertor.ToString(meta))
		server.meta[stream.ID()] = &meta

		log.Println(len(server.conn))
	})
}

func (server *YunServer) Addr() string {
	return (*server.h).Addrs()[0].Encapsulate(multiaddr.StringCast("/p2p/" + (*server.h).ID().String())).String()
}

func New(ctx context.Context, h *host.Host) *YunServer {
	s := &YunServer{
		ctx,
		h,
		map[string]*jsonrpc2.Conn{},
		map[string]*common.Meta{},
	}
	s.initS()

	return s
}
