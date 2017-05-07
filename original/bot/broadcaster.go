package bot

import (
	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

type (
	Broadcaster struct {
		BotIn chan *Bot
		bots  map[*Bot]bool
		msgIn chan *model.Message
	}
)

func (b *Broadcaster) Run() {
	for {
		select {
		case bot := <-b.BotIn:
			b.bots[bot] = true
		case msg := <-b.msgIn:
			for bot, _ := range b.bots {
				bot.in <- msg
			}
		}
	}
}

func NewBroadcaster(msgIn chan *model.Message) *Broadcaster {
	memberIn := make(chan *Bot)
	return &Broadcaster{
		BotIn: memberIn,
		bots:  make(map[*Bot]bool),
		msgIn: msgIn,
	}
}
