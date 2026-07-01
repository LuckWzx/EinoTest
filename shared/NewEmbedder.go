package shared

import (
	"context"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"os"
)

// NewEmbedder 创建 ARK 多模态 embedding 实例
func NewEmbedder(ctx context.Context) *ark.Embedder {
	//_ = godotenv.Load() // 幂等加载，已加载则跳过
	apiType := ark.APITypeMultiModal
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("EMBEDDER"),
		APIType: &apiType,
	})

	if err != nil {
		panic(err)
	}
	return embedder
}
