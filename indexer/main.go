package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
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
	InitClient()
	ctx := context.Background()

	apiType := ark.APITypeMultiModal
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("EMBEDDER"),
		APIType: &apiType,
	})

	if err != nil {
		panic(err)
	}

	var collection = "test"

	// 删除旧 collection（schema 变更后必须重建）
	_ = MilvusCli.DropCollection(ctx, collection)

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
			DataType: entity.FieldTypeFloatVector,
			TypeParams: map[string]string{
				"dim": "2048",
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
		Client:            MilvusCli,
		Collection:        collection,
		Fields:            fields,
		Embedding:         embedder,
		DocumentConverter: floatDocumentConverter,
		MetricType:        milvus.COSINE,
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
	}

	ids, err := indexer.Store(ctx, docs)
	if err != nil {
		panic(err)
	}
	fmt.Println(ids)

}

// floatDocRow 自定义文档行结构，Vector 为 []float32 适配 FloatVector 字段
type floatDocRow struct {
	ID       string    `milvus:"name:id"`
	Content  string    `milvus:"name:content"`
	Vector   []float32 `milvus:"name:vector"`
	Metadata []byte    `milvus:"name:metadata"`
}

func floatDocumentConverter(ctx context.Context, docs []*schema.Document, vectors [][]float64) ([]interface{}, error) {
	rows := make([]interface{}, 0, len(docs))
	for i, doc := range docs {
		// float64 -> float32
		float32Vec := make([]float32, len(vectors[i]))
		for j, v := range vectors[i] {
			float32Vec[j] = float32(v)
		}
		// metadata -> JSON bytes
		metaBytes, err := json.Marshal(doc.MetaData)
		if err != nil {
			return nil, fmt.Errorf("marshal metadata: %w", err)
		}
		rows = append(rows, &floatDocRow{
			ID:       doc.ID,
			Content:  doc.Content,
			Vector:   float32Vec,
			Metadata: metaBytes,
		})
	}
	return rows, nil
}
