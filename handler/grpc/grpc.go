package grpc

import (
	"github.com/dailoi280702/vrs-general-service/proto"
	"google.golang.org/grpc"
)

type Service struct {
	proto.UnimplementedServiceServer
}

func NewGRPCServer() *grpc.Server {
	return grpc.NewServer()
}
