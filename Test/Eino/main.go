package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/schema"
)
type Mess struct {
    Name string `json:"name"`
    Num int `json:"num"`
}


func main() {
    ctx := context.Background()
    
	model, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: "http://localhost:11434", // Ollama 服务地址
		Model:   "llama2",                 // 模型名称
	})
	
    if err != nil {
        panic(err)
    }
    
    // 准备消息
    messages := []*schema.Message{
        // schema.SystemMessage("请你模仿一个电商客服,现在用户会发送消息让你帮助他下单,请你分析理解用户的消息,并且将输出一个json数据,其中包括两个关键字:商品种类、商品数量,两个关键字对应于两个列表数据,支持用户下单多种商品,但要保证两个列表中的数据一一对应."),
        schema.UserMessage("请提取以下句子中的名词和量词,只返回一个json格式数据,名词的关键字设为name,量词的关键字设为num:下单一箱牛奶"),
    }

    // 生成回复
    response, err := model.Generate(ctx, messages)
    if err != nil {
        panic(err)
    }

    var mess Mess
    err = json.Unmarshal([]byte(response.Content), &mess)
    if err != nil {
        fmt.Println("json wrong:", err)
        return
    }
    
    println(mess.Name)
    println(mess.Num)

    // 处理回复
    println(1)
    println(response.Content)
}
