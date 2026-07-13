package main

import (
	"EinoTest/shared"
	"context"
	"fmt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()
	//注册图
	g := compose.NewGraph[map[string]string, *schema.Message]()
	//编写节点
	lambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output map[string]string, err error) {
		if input["role"] == "tsunderes" {
			return map[string]string{"role": "傲娇", "content": input["content"]}, nil
		}
		if input["role"] == "cute" {
			return map[string]string{"role": "可爱", "content": input["content"]}, nil
		}
		return map[string]string{"role": "", "content": input["content"]}, nil
	})
	TsundereLambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output []*schema.Message, err error) {
		return []*schema.Message{
			{
				Role:    schema.System,
				Content: "你是一个高冷傲娇的大小姐，每次都会用傲娇的语气回答我的问题",
			},
			{
				Role:    schema.User,
				Content: input["content"],
			},
		}, nil
	})
	CuteLambda := compose.InvokableLambda(func(ctx context.Context, input map[string]string) (output []*schema.Message, err error) {
		return []*schema.Message{
			{
				Role:    schema.System,
				Content: "你是一个可爱的小女孩，每次都会用可爱的语气回答我的问题",
			},
			{
				Role:    schema.User,
				Content: input["content"],
			},
		}, nil
	})

	model := shared.NewArkChatModel(ctx)

	if err != nil {
		panic(err)
	}

	//加入节点
	err = g.AddLambdaNode("lambda", lambda)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("tsundere", TsundereLambda)
	if err != nil {
		panic(err)
	}
	err = g.AddLambdaNode("cute", CuteLambda)
	if err != nil {
		panic(err)
	}
	err = g.AddChatModelNode("model", model)
	if err != nil {
		panic(err)
	}
	//编写分支
	g.AddBranch("lambda", compose.NewGraphBranch(func(ctx context.Context, in map[string]string) (endNode string, err error) {
		if in["role"] == "傲娇" {
			return "tsundere", nil
		}
		if in["role"] == "可爱" {
			return "cute", nil
		}
		return "tsundere", nil
	}, map[string]bool{"tsundere": true, "cute": true}))
	//节点连接
	err = g.AddEdge(compose.START, "lambda")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("tsundere", "model")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("cute", "model")
	if err != nil {
		panic(err)
	}
	err = g.AddEdge("model", compose.END)
	if err != nil {
		panic(err)
	}
	//编译运行
	r, err := g.Compile(ctx)
	if err != nil {
		panic(err)
	}
	input := map[string]string{
		"role":    "tsundere",
		"content": "你好啊，现在几岁了",
	}
	answer, err := r.Invoke(ctx, input)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer.Content)
}
