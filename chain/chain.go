package main

import (
	"EinoTest/shared"
	"context"
	"fmt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

// 链式编排
func main() {
	ctx := context.Background()

	model := shared.NewArkChatModel(ctx)

	lambda := compose.InvokableLambda(func(ctx context.Context, input string) (output []*schema.Message, err error) {
		desuwa := input + "回答结尾加上desuwa"
		output = []*schema.Message{
			{
				Role:    schema.User,
				Content: desuwa,
			},
		}
		return output, nil
	})

	chain := compose.NewChain[string, *schema.Message]()
	chain.AppendLambda(lambda).AppendChatModel(model)
	r, err := chain.Compile(ctx)
	if err != nil {
		panic(err)
	}
	answer, err := r.Invoke(ctx, "你好，你可以告诉我你的名字吗？")
	if err != nil {
		panic(err)
	}
	fmt.Println(answer.Content)

}
