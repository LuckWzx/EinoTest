package shared

import (
	"context"
	"fmt"
	cli "github.com/milvus-io/milvus-sdk-go/v2/client"
)

var MilvusCli cli.Client

// InitClient 初始化 Milvus 客户端
func InitClient() {
	ctx := context.Background()
	client, err := cli.NewClient(ctx, cli.Config{
		Address: "39.105.40.22:19530",
		DBName:  "WZXEinoFrame",
	})
	if err != nil {
		panic(fmt.Errorf("failed to create milvus client: %w", err))
	}
	MilvusCli = client
}
