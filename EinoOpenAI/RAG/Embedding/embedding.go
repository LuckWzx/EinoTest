package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/libs/acl/openai"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load("./EinoOpenAI/.env")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	httpCli := &http.Client{
		Timeout: 60 * time.Second,
	}
	embedder, err := openai.NewEmbeddingClient(ctx, &openai.EmbeddingConfig{
		APIKey:     os.Getenv("QWEN_API_KEY"),
		Model:      os.Getenv("QWEN_EMBED_MODEL"),
		BaseURL:    os.Getenv("QWEN_BASE_URL"),
		HTTPClient: httpCli,
	})
	if err != nil {
		panic(err)
	}
	input := []string{
		"你好，拟好",
	}
	embeddings, err := embedder.EmbedStrings(ctx, input)
	if err != nil {
		panic(err)
	}
	//fmt.Print(embeddings)
	for _, embedding := range embeddings {
		fmt.Println(len(embedding))
	}
	//fmt.Println(len(embeddings))

}
