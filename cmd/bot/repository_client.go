package main

import (
	"github.com/pkg/errors"
	pb "gitlab.ozon.dev/Woofka/movie-review-system/pkg/repository_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitRepositoryClient() (pb.RepositoryClient, error) {
	conns, err := grpc.Dial(":8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "InitRepositoryClient Dial error")
	}
	client := pb.NewRepositoryClient(conns)
	return client, nil
}
