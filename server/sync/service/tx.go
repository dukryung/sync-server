package service

import (
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"sync-server/server/types"
)

type TxService struct {
	ctx      types.Context
	config   types.SyncServer
	tmClient tmservice.ServiceClient
	txClient tx.ServiceClient
}

func NewTxService(ctx types.Context, config types.SyncServer) *TxService {
	return &TxService{
		ctx:    ctx,
		tmClient: tmservice.NewServiceClient(ctx.GRPCConn()),
		txClient: tx.NewServiceClient(ctx.GRPCConn()),


		config: config,
	}
}
