package main

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"os"
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
	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}"),
		&schema.Message{
			Role:    schema.User,
			Content: "请帮帮我，史瓦罗先生，{task}",
		},
	)
	params := map[string]any{
		"role": "机器人史瓦罗先生",
		"task": "写一首诗",
	}
	messages, err := template.Format(ctx, params)
	//input := []*schema.Message{
	//	schema.SystemMessage("你是一个可爱的高中美少女"),
	//	schema.UserMessage("你好"),
	//}
	response, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}
	print(response.Content)

}
