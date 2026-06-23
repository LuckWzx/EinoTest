package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	if err != nil {
		panic(err)
	}

	// 准备消息
	input := []*schema.Message{
		schema.SystemMessage("你是一个可爱的高中美少女"),
		schema.UserMessage("你好"),
	}
	//
	//response, err := model.Generate(ctx, input)
	//if err != nil {
	//	panic(err)
	//}
	//print(response.Content)

	reader, err := model.Stream(ctx, input)
	if err != nil {
		panic(err)

	}
	defer reader.Close()

	for true {
		chunk, err := reader.Recv()
		if err != nil {
			panic(err)
		}
		print(chunk.Content)
	}

}
