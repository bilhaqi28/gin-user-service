package main

import (
	"context"
	"errors"
	"gin-user-service/config"
	"gin-user-service/repository"
	pb "gin-user-service/servergrpc/model/user"
	"gin-user-service/service"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type dataUserServer struct {
	pb.UnimplementedDataUserServer
	mu      sync.Mutex
	service service.ServiceUser
}

func (d *dataUserServer) GenerateJwtByToken(ctx context.Context, tokenGrand *pb.TokenGrand) (*pb.TokenJwt, error) {
	var tokenJwt pb.TokenJwt
	result, err := d.service.ShowByToken(ctx, tokenGrand.Token)
	if err != nil {
		return &tokenJwt, errors.New(err.Error())
	}
	tokenJwt.Token = result.TokenJwt
	return &tokenJwt, nil
}

func newUserServer() *dataUserServer {
	db := config.NewDB()
	repository := repository.NewRepositoryUser()
	service := service.NewServiceUser(repository, db)
	return &dataUserServer{
		service: service,
	}

}

func main() {
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("Error in lister", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDataUserServer(grpcServer, newUserServer())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("Error Running Server GRPC", err.Error())
	}
}
