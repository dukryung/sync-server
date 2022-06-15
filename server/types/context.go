package types

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"google.golang.org/grpc"
)

type Context struct {
	ctx        context.Context
	grpcClient *grpc.ClientConn
	cancel     context.CancelFunc
	txConfig   client.TxConfig

}

func (c Context) GRPCConn() *grpc.ClientConn { return c.grpcClient }

func NewContext() Context {
	return Context{
		ctx: context.Background(),
	}
}

func NewContextCancel() Context {
	ctx, cancel := context.WithCancel(context.Background())
	return Context{
		ctx:    ctx,
		cancel: cancel,
		txConfig: klaatoo.MakeEncoding
	}
}

func (c Context) WithGRPCConn(conn *grpc.ClientConn) Context {
	if conn == nil {

	}
	c.grpcClient = conn
	return c
}
