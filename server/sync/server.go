package sync

import (
	"sync-server/server/types"
)

type Server struct {
	ctx types.Context
	close chan bool
	txService *service.T
}

func NewServer(config *types.SyncServer) *Server {
	s := Server{}

	grpcConn, err := config.GetGRPCConnection()
	if err != nil {

	}

	s.ctx = types.NewContextCancel().WithGRPCConn(grpcConn)

	return &Server{}
}

func (s *Server) Run() {

}

func (s *Server) Close() {

}
