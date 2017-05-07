package bot

import "github.com/VG-Tech-Dojo/vg-1day-2017/original/model"

type (
	// messageを受け取り、投稿用messageを作るインターフェース
	Processor interface {
		Process(message *model.Message) *model.Message
	}

	// "hello, world!"メッセージを作るprocessor
	HelloWorldProcessor struct {}

	// "大吉", "吉", "中吉", "小吉", "末吉", "凶"のいずれかをランダムで作るprocessor
	OmikujiProcessor struct {}
)

// "hello, world!"メッセージを作る
func (p *HelloWorldProcessor) Process(msgIn *model.Message) *model.Message {
	return &model.Message{
		Body: msgIn.Body + ", world!",
	}
}

// "大吉", "吉", "中吉", "小吉", "末吉", "凶"のいずれかをランダムで作る
func (p *OmikujiProcessor) Process(msgIn *model.Message) *model.Message {
	fortunes := []string{
		"大吉",
		"吉",
		"中吉",
		"小吉",
		"末吉",
		"凶",
	}
    result := fortunes[randIntn(len(fortunes))]
	return &model.Message{
		Body: result,
	}
}
