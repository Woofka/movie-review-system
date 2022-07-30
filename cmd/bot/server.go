package main

import (
	"net"

	apiPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/api"
	reviewPkg "gitlab.ozon.dev/Woofka/movie-review-system/internal/pkg/core/review"
	pb "gitlab.ozon.dev/Woofka/movie-review-system/pkg/api"
	"google.golang.org/grpc"
)

func runGRPCServer(review reviewPkg.Interface) {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New(review))

	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
