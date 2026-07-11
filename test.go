package main

//
//import (
//	"context"
//	"io"
//	"testing"
//
//	"github.com/cloudwego/eino/compose"
//	"github.com/cloudwego/eino/schema"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestTypeMatch(t *testing.T) {
//	ctx := context.Background()
//
//	g1 := compose.NewGraph[[]*schema.Message, string]()
//	_ = g1.AddChatModelNode("model", &mockChatModel{})
//	_ = g1.AddLambdaNode("lambda", compose.InvokableLambda(func(_ context.Context, msg *schema.Message) (string, error) {
//		return msg.Content, nil
//	}))
//	_ = g1.AddEdge(compose.START, "model")
//	_ = g1.AddEdge("model", "lambda")
//	_ = g1.AddEdge("lambda", compose.END)
//
//	runner, err := g1.Compile(ctx)
//	assert.NoError(t, err)
//
//	c, err := runner.Invoke(ctx, []*schema.Message{
//		schema.UserMessage("what's the weather in beijing?"),
//	})
//	assert.NoError(t, err)
//	assert.Equal(t, "the weather is good", c)
//
//	s, err := runner.Stream(ctx, []*schema.Message{
//		schema.UserMessage("what's the weather in beijing?"),
//	})
//	assert.NoError(t, err)
//
//	var fullStr string
//	for {
//		chunk, err := s.Recv()
//		if err != nil {
//			if err == io.EOF {
//				break
//			}
//			panic(err)
//		}
//
//		fullStr += chunk
//	}
//	assert.Equal(t, c, fullStr)
//}
