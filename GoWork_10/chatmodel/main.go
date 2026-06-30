package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load("./GoWork_10/.env")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	model, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  os.Getenv("API_KEY"),
		Model:   os.Getenv("DEEPSEEK_CHAT_MODEL"),
		BaseURL: os.Getenv("BASE_URL"),
	})
	template := prompt.FromMessages(
		schema.FString,
		schema.SystemMessage("你是一个{role}"),
		&schema.Message{
			Role:    schema.User,
			Content: "请帮帮我，史瓦罗先生，帮我解决{task}",
		},
	)
	params := map[string]any{
		"role": "机器人史瓦罗先生",
		"task": "写一首诗",
	}
	//input := []*schema.Message{
	//	schema.SystemMessage("你是一个可爱的高中美少女"),
	//	schema.UserMessage("你好"),
	//}

	//全量输出
	//chat prompt template使用
	message, err := template.Format(ctx, params)
	response, err := model.Generate(ctx, message)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.Content)

	//流式输出
	//reader, err := model.Stream(ctx, input)
	//if err != nil {
	//	panic(err)
	//}
	//defer reader.Close()
	//for {
	//	chunk, err := reader.Recv()
	//	if err != nil {
	//		if errors.Is(err, io.EOF) {
	//			break
	//		}
	//		panic(err)
	//	}
	//	fmt.Print(chunk.Content)
	//}

}
