package shared

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/joho/godotenv"
	"os"
)

// NewArkChatModel NewArkModel 初始化chatmodel
func NewArkChatModel(ctx context.Context) *ark.ChatModel {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	if err != nil {
		panic(err)
	}

	return model
}
