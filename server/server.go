package server

import (
	"context"
	"log"

	"Github.com/Ekaloi/library/api"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Server struct {
	client *api.Client
	dynamo *dynamodb.Client
}


func NewServer() *Server{
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}
	return &Server{client: api.NewClient(), dynamo: dynamodb.NewFromConfig(cfg)}
}
