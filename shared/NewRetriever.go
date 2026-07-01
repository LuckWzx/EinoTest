package shared

import (
	"context"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/retriever/milvus"
)

func NewArkRetriever(ctx context.Context, embedder *ark.Embedder) *milvus.Retriever {

	retriever, err := milvus.NewRetriever(ctx, &milvus.RetrieverConfig{
		Client:      MilvusCli,
		Collection:  "test",
		VectorField: "vector",
		OutputFields: []string{
			"id",
			"content",
			"metadata",
		},
		TopK:      2, // 召回文档上限
		Embedding: embedder,
	})
	if err != nil {
		panic(err)
	}

	return retriever

}
