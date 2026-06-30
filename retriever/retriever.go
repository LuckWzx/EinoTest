package main

import (
	"context"
	"fmt"

	"EinoTest/shared"

	"github.com/cloudwego/eino-ext/components/retriever/milvus"
)

func main() {
	ctx := context.Background()
	Retriever(ctx)

}

func Retriever(ctx context.Context) {
	shared.InitClient()
	embedder, err := shared.NewEmbedder(ctx)
	if err != nil {
		panic(err)
	}

	retriever, err := milvus.NewRetriever(ctx, &milvus.RetrieverConfig{
		Client:      shared.MilvusCli,
		Collection:  "test",
		Partition:   nil,
		VectorField: "vector",
		OutputFields: []string{
			"id",
			"vector",
			"content",
			"metadata",
		},
		TopK:      1, // 召回文档上限
		Embedding: embedder,
	})
	if err != nil {
		panic(err)
	}

	results, err := retriever.Retrieve(ctx, "魏正想")
	if err != nil {
		panic(err)
	}
	fmt.Println(results[0].Content)
}
