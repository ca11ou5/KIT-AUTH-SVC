package main

import (
	"Auth_Service/configs"
	"Auth_Service/internal/client"
	"Auth_Service/internal/pb"
	"Auth_Service/internal/repository"
	"Auth_Service/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := configs.InitConfig()

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatal("failed to listen " + cfg.Port + ": " + err.Error())
	}

	//GRPC Server
	client1 := client.InitSmsClient(cfg.SmsSvcUrl)
	grpcServer := grpc.NewServer()
	s := service.Server{
		Repo:   repository.InitRepository(&cfg),
		Client: client1,
	}
	pb.RegisterAuthServiceServer(grpcServer, &s)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("failed to serve grpc server: " + err.Error())
	}
}
