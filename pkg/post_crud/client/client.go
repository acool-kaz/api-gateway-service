package client

import (
	"fmt"
	"log"

	post_crud_pb "github.com/acool-kaz/api-gateway-service/pkg/post_crud/pb"
	"google.golang.org/grpc"
)

type PostCRUDClientConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func InitPostCRUDClientConfig(host, port string) *PostCRUDClientConfig {
	log.Println("init post crud client config")

	return &PostCRUDClientConfig{
		Host: host,
		Port: port,
	}
}

type PostCRUDClient struct {
	client post_crud_pb.PostCRUDServiceClient
}

func InitPostCRUDClient(cfg *PostCRUDClientConfig) (*PostCRUDClient, error) {
	log.Println("init post crud client")

	conn, err := grpc.Dial(cfg.Host+":"+cfg.Port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("init post crud client: %w", err)
	}

	return &PostCRUDClient{
		client: post_crud_pb.NewPostCRUDServiceClient(conn),
	}, nil
}
