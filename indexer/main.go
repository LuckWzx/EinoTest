package main

import (
	"context"
	"fmt"

	"EinoTest/shared"

	"github.com/cloudwego/eino-ext/components/indexer/milvus"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	//初始化客户端
	shared.InitClient()
	ctx := context.Background()

	embedder := shared.NewEmbedder(ctx)
	if err != nil {
		panic(err)
	}

	var collection = "test"

	// 删除旧 collection（schema 变更后必须重建）
	//_ = shared.MilvusCli.DropCollection(ctx, collection)

	var fields = []*entity.Field{
		{
			Name:     "id",
			DataType: entity.FieldTypeVarChar,
			TypeParams: map[string]string{
				"max_length": "256",
			},
			PrimaryKey: true,
		},
		{
			Name:     "vector", // 确保字段名匹配
			DataType: entity.FieldTypeBinaryVector,
			TypeParams: map[string]string{
				"dim": "65536",
			},
		},
		{
			Name:     "content",
			DataType: entity.FieldTypeVarChar,
			TypeParams: map[string]string{
				"max_length": "8192",
			},
		},
		{
			Name:     "metadata",
			DataType: entity.FieldTypeJSON,
		},
	}

	indexer, err := milvus.NewIndexer(ctx, &milvus.IndexerConfig{
		Client:     shared.MilvusCli,
		Collection: collection,
		Fields:     fields,
		Embedding:  embedder,
	})
	if err != nil {
		panic(err)
	}

	docs := []*schema.Document{
		{
			ID:      "1",
			Content: "你说得对。但是原神是一款二次元开放大世界游戏",
			MetaData: map[string]any{
				"author": "木乔",
			},
		},

		{
			ID:      "2",
			Content: "你说得对。但是原神是一款二次元开放大世界游戏",
			MetaData: map[string]any{
				"author": "鹰角",
			},
		},
		{
			ID:      "3",
			Content: "魏正想是软件学院的学生",
			MetaData: map[string]any{
				"author": "魏正想",
			},
		},

		{
			ID:      "4",
			Content: "王者荣耀3天前刚刚更新",
			MetaData: map[string]any{
				"author": "想想",
			},
		},
	}

	ids, err := indexer.Store(ctx, docs)
	if err != nil {
		panic(err)
	}
	fmt.Println(ids)

}
