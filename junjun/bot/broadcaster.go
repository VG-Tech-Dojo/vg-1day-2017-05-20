package bot

import (
	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/junjun/model"
)

type (
	// Broadcaster は1つのチャンネルで複数botを動かすためのヘルパーです
	//
	// msgInで受け取ったmessageをbotsに登録された全botに渡します
	//
	// botsへの登録はBotInで行います
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

// Run はBroadcasterを起動します
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

// NewBroadcaster は新しいBroadcaster構造体のポインタを返します
func NewBroadcaster(msgIn chan *model.Message) *Broadcaster {
	memberIn := make(chan *Bot)
	return &Broadcaster{
		BotIn: memberIn,
		bots:  []*Bot{},
		msgIn: msgIn,
	}
}
