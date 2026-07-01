package main

import (
	"EinoTest/shared"
	"context"
	"fmt"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	//初始化chatmodel
	//model := shared.NewArkChatModel(ctx)

	//初始化milvusClient
	shared.InitClient()

	//初始化Embedder
	embedder := shared.NewEmbedder(ctx)

	//初始化Indexer
	indexer := shared.NewArkIndexer(ctx, embedder)

	//初始化retriever
	//retriever := shared.NewArkRetriever(ctx, embedder)

	//初始化Transformer
	transformer := shared.NewTransformer(ctx)

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
	// 转换文档
	//transformedDocs, _ := transformer.Transform(ctx, []*schema.Document{markdownDoc})
	transformedDocs, _ := transformer.Transform(ctx, docs)

	//for i, doc := range transformedDocs {
	//	println("片段", i+1, ":", doc.Content)
	//	println("标题层级：")
	//	for k, v := range doc.MetaData {
	//		if k == "h1" || k == "h2" || k == "h3" {
	//			println("  ", k, ":", v)
	//		}
	//	}
	//}
	//
	//fmt.Println("=====================================================================")
	//fmt.Println(len(docs))
	for i, doc := range transformedDocs {

		doc.ID = docs[0].ID + "_" + strconv.Itoa(i)
		println(doc.ID)

	}

	ids, err := indexer.Store(ctx, transformedDocs)
	if err != nil {
		panic(err)
	}
	fmt.Println(ids)

}
