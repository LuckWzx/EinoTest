package main

import (
	"context"
	cli "github.com/milvus-io/milvus-sdk-go/v2/client"
	"log"
)

var MilvusCli cli.Client

func InitClient() {
	ctx := context.Background()
	client, err := cli.NewClient(ctx, cli.Config{
		Address: "localhost:19530",
		DBName:  "Eino",
	})
	if err != nil {
		log.Fatalf("failed to create client:%v", err)
	}
	MilvusCli = client
	log.Println("milvus client init success")
}
