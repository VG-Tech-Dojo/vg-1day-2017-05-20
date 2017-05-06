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
)

func (p *HelloWorldProcessor) Process(msgIn *model.Message) *model.Message {
	return &model.Message{
		Body: msgIn.Body + ", world!",
	}
}
