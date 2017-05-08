package bot

import (
	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

type (
	// 1つのチャンネルで複数botを動かすためのヘルパー
	//
	// msgInで受け取ったmessageをbotsに登録された全botに渡す
	//
	// botsへの登録はBotInで行う
	//
	//   fields
	// 	   BotIn chan *Bot
	// 	   bots  map[*Bot]bool
	// 	   msgIn chan *model.Message
	Broadcaster struct {
		BotIn chan *Bot
		bots  []*Bot
		msgIn chan *model.Message
	}
)

// broadcasterを起動する
func (b *Broadcaster) Run() {
	for {
		select {
		case bot := <-b.BotIn:
			b.bots = append(b.bots, bot)
		case msg := <-b.msgIn:
			for _, bot := range b.bots {
				bot.in <- msg
			}
		}
	}
}

// broadcasterのインスタンス生成はこの関数を使う
func NewBroadcaster(msgIn chan *model.Message) *Broadcaster {
	memberIn := make(chan *Bot)
	return &Broadcaster{
		BotIn: memberIn,
		bots:  []*Bot{},
		msgIn: msgIn,
	}
}
