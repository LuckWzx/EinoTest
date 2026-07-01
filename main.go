package main

import (
	"EinoTest/shared"
	"context"
	"github.com/joho/godotenv"
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
	//indexer := shared.NewArkIndexer(ctx, embedder)

	//初始化retriever
	retriever := shared.NewArkRetriever(ctx, embedder)
	results, err := retriever.Retrieve(ctx, "如“三人成虎”“聒噪如雷”“聪明一世”等")
	if err != nil {
		panic(err)

	}
	for _, doc := range results {
		println(doc.ID)
		println(doc.Content)
		println("==================================================")

	}

	//初始化Transformer
	//transformer := shared.NewTransformer(ctx)
	//
	//content, err := os.OpenFile("./transformer/document.md", os.O_CREATE|os.O_RDWR, 0755)
	//if err != nil {
	//	panic(err)
	//}
	//defer content.Close()
	//
	//bs, err := os.ReadFile("./transformer/document.md")
	//if err != nil {
	//	panic(err)
	//}

	//数据切分
	//docs := []*schema.Document{
	//	{
	//		ID:      "doc1",
	//		Content: string(bs),
	//	},
	//}
	// 转换文档
	//transformedDocs, _ := transformer.Transform(ctx, []*schema.Document{markdownDoc})

	//transformedDocs, _ := transformer.Transform(ctx, docs)
	//for i, doc := range transformedDocs {
	//
	//	doc.ID = docs[0].ID + "_" + strconv.Itoa(i)
	//	println(doc.ID)
	//
	//}
	//
	//ids, err := indexer.Store(ctx, transformedDocs)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ids)

}
