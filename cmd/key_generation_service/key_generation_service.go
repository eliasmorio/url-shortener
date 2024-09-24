package main

import (
	"UrlShortener/internal/logging"
	"UrlShortener/internal/service"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

func main() {
	logging.Init()

	slog.Info("Starting key generation service")
	lis, err := net.Listen("tcp", ":8081")

	grpcServer := grpc.NewServer()

	service.RegisterKgsServer(grpcServer, service.NewKeygenServer())

	err = grpcServer.Serve(lis)

	if err != nil {
		slog.Error("failed to serve", err)
		panic(err)
	}
}
