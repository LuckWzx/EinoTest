package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	cli "github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

var MilvusCli cli.Client

// InitClient 初始化 Milvus 客户端
func InitClient() {
	ctx := context.Background()
	client, err := cli.NewClient(ctx, cli.Config{
		Address: "localhost:19530",
		DBName:  "AwesomeEino",
	})
	if err != nil {
		panic(fmt.Errorf("failed to create milvus client: %w", err))
	}
	MilvusCli = client
}

// NewEmbedder 创建 ARK 多模态 embedding 实例
func NewEmbedder(ctx context.Context) (*ark.Embedder, error) {
	_ = godotenv.Load() // 幂等加载，已加载则跳过
	apiType := ark.APITypeMultiModal
	return ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("EMBEDDER"),
		APIType: &apiType,
	})
}

// FloatDocRow 自定义文档行结构，Vector 为 []float32 适配 FloatVector 字段
type FloatDocRow struct {
	ID       string    `milvus:"name:id"`
	Content  string    `milvus:"name:content"`
	Vector   []float32 `milvus:"name:vector"`
	Metadata []byte    `milvus:"name:metadata"`
}

// FloatVectorConverter 将 float64 向量转为 FloatVector（而非默认的 BinaryVector）
func FloatVectorConverter(ctx context.Context, vectors [][]float64) ([]entity.Vector, error) {
	vec := make([]entity.Vector, 0, len(vectors))
	for _, v := range vectors {
		float32Vec := make([]float32, len(v))
		for j, val := range v {
			float32Vec[j] = float32(val)
		}
		vec = append(vec, entity.FloatVector(float32Vec))
	}
	return vec, nil
}
func FloatDocumentConverter(ctx context.Context, docs []*schema.Document, vectors [][]float64) ([]interface{}, error) {
	rows := make([]interface{}, 0, len(docs))
	for i, doc := range docs {
		float32Vec := make([]float32, len(vectors[i]))
		for j, v := range vectors[i] {
			float32Vec[j] = float32(v)
		}
		metaBytes, err := json.Marshal(doc.MetaData)
		if err != nil {
			return nil, fmt.Errorf("marshal metadata: %w", err)
		}
		rows = append(rows, &FloatDocRow{
			ID:       doc.ID,
			Content:  doc.Content,
			Vector:   float32Vec,
			Metadata: metaBytes,
		})
	}
	return rows, nil
}
