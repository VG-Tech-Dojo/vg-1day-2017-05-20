package bot

import "github.com/VG-Tech-Dojo/vg-1day-2017/original/model"

type (
	// Processor
	// messageを受け取り、投稿用のmessageを作成するインターフェース
	Processor interface {
		Process(message *model.Message) *model.Message
	}

	// HelloWorldProcessor
	// hello, world!メッセージを作成するProcessor
	HelloWorldProcessor struct {}

	// OmikujiProcessor
	// 大吉、吉、中吉、小吉、末吉、凶のいずれかをランダムで返すProcessor
	OmikujiProcessor struct {}
)

func (p *HelloWorldProcessor) Process(msgIn *model.Message) *model.Message {
	return &model.Message{
		Body: msgIn.Body + ", world!",
	}
}

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
