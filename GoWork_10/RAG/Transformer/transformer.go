package main

import (
	"context"
	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	splitter, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		Headers: map[string]string{
			"#":    "h1",
			"##":   "h2",
			"####": "h3",
		},
		TrimHeaders: false,
	})
	if err != nil {
		panic(err)
	}
	content, err := os.OpenFile("./Transformer/document.md", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer content.Close()
	bs, err := os.ReadFile("./Transformer/document.md")
	if err != nil {
		panic(err)
	}
	if len(bs) == 0 {
		log.Fatal("document.md 文件内容为空")
	}
	docs := []*schema.Document{
		{
			ID:      "doc1",
			Content: string(bs),
		},
	}
	results, err := splitter.Transform(ctx, docs)
	if err != nil {
		panic(err)
	}
	for i, doc := range results {
		println("片段", i+1, ":", doc.Content)
		println("标题层级：")
		for k, v := range doc.MetaData {
			if k == "h1" || k == "h2" || k == "h3" {
				println("  ", k, ":", v)
			}
		}
	}
}
