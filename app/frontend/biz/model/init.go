package model

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino-ext/components/model/ollama"
)

var (
	Model *ollama.ChatModel
	Err error
)

func Init() {
	ctx := context.Background()

	Model, Err = MyNewChatModel(ctx)
}

func MyNewChatModel(ctx context.Context) (*ollama.ChatModel, error){

	m, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
	BaseURL: "http://localhost:11434", // Ollama 服务地址
	Model:   "llama2",                 // 模型名称
	})
	fmt.Println("1")
	if err != nil {
		panic(err)
	}
	return m, nil
}
