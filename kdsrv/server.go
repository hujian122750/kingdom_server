package kdsrv

import (
	"context"
	pb "github.com/bench/kingdom-server/proto/pb"
)

var RpcServer *Server

func init() {
	RpcServer = &Server{}
}

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) GetUser(context.Context, *pb.UserRequest) (*pb.UserResponse, error) {
	rsp := &pb.UserResponse{
		Name:  "hahah",
		Age:   10,
		Email: "11111@163.com",
	}
	return rsp, nil
}
