package gapi

import (
	"github.com/aldisaputra17/go-micro/src/domain/pb"
	"github.com/aldisaputra17/go-micro/src/module"
)

type GRPCServer struct {
	pb.UnimplementedUserServiceServer
	Address string
	mdl     *module.Module
}

func NewGRPCServer(address string, mdl *module.Module) (gapi *GRPCServer, err error) {
	return &GRPCServer{
		Address: address,
		mdl:     mdl,
	}, nil
}
