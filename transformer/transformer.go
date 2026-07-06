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
	// 初始化 transformer (以 markdown 为例)
	transformer, _ := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		// 配置参数
		Headers: map[string]string{
			"#":   "h1",
			"##":  "h2",
			"###": "h3",
		},
		TrimHeaders: false,
	})

	content, err := os.OpenFile("./transformer/document.md", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		panic(err)

	}
	defer content.Close()

	bs, err := os.ReadFile("./transformer/document.md")
	if err != nil {
		panic(err)
	}
	docs := []*schema.Document{
		{
			ID:      "doc1",
			Content: string(bs),
		},
	}

	//markdownDoc := &schema.Document{
	//	Content: "## Title 1\nHello Word\n## Title 2\nWord Hello",
	//}
	// 转换文档
	//transformedDocs, _ := transformer.Transform(ctx, []*schema.Document{markdownDoc})
	transformedDocs, _ := transformer.Transform(ctx, docs)

	for idx, doc := range transformedDocs {
		log.Printf("doc segment %v: %v", idx, doc.Content)
	}
	// 处理分割结果
	for i, doc := range transformedDocs {
		println("片段", i+1, ":", doc.Content)
		println("标题层级：")
		for k, v := range doc.MetaData {
			if k == "h1" || k == "h2" || k == "h3" {
				println("  ", k, ":", v)
			}
		}
	}

}

//markdownDoc := &schema.Document{
//	Content: "## Title 1\nHello Word\n## Title 2\nWord Hello",
//}
