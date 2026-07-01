package shared

import (
	"context"
	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino/components/document"
)

func NewTransformer(ctx context.Context) document.Transformer {
	// 初始化 transformer (以 markdown 为例)
	transformer, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		// 配置参数
		Headers: map[string]string{
			"#":   "h1",
			"##":  "h2",
			"###": "h3",
		},
		TrimHeaders: false,
	})

	if err != nil {
		panic(err)
	}

	return transformer
}
