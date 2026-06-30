package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/tool/browseruse"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	but, err := browseruse.NewBrowserUseTool(ctx, &browseruse.Config{})
	if err != nil {
		panic(err)
	}
	url := "https://bilibili.com"
	resutl, err := but.Execute(&browseruse.Param{
		Action: browseruse.ActionGoToURL,
		URL:    &url,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(resutl)
	time.Sleep(10 * time.Second)
	but.Cleanup()
}
