package types

import (
	"encoding/json"
	"google.golang.org/grpc"
	"os"
)
const (
	ConfigPath  = "./config.json"
)

type AppConfig struct {
	Sync *SyncServer
}

type SyncServer struct {
	Enable bool   `json:"enable"`
	Node   string `json:"node"`
	GRPC   string `json:"grpc"`
}



func (config *AppConfig) LoadConfig(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		return err
	}

	return nil
}

func (config *SyncServer) GetGRPCConnection() (*grpc.ClientConn,error) {
	return grpc.Dial(config.GRPC,grpc.WithInsecure())
}
