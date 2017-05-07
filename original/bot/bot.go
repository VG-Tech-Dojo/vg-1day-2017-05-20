package bot

import (
	"context"
	"fmt"

	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

type (
	// inで受け取ったmessageがcheckerの条件を満たした場合、processorが投稿用messageを作り、outに渡す
	//
	//   fields
	//     name      string
	//     in        chan *model.Message
	//     out       chan *model.Message
	//     checker   Checker
	//     processor Processor
	Bot struct {
		name      string
		in        chan *model.Message
		out       chan *model.Message
		checker   Checker
		processor Processor
	}
)

// botを起動する
func (b *Bot) Run(ctx context.Context) {
	fmt.Printf("%s start\n", b.name)

	// メッセージ監視
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s stop", b.name)
			return
		case m := <-b.in:
			fmt.Printf("%s received: %v\n", b.name, m)

			if b.checker.Check(m) {
				b.respond(m)
			}
		}
	}
}

// "hello"を受け取ると"hello, world!"を返すbot
func NewSimpleBot(out chan *model.Message) *Bot {
	in := make(chan *model.Message)

	checker := NewRegexpChecker("\\Ahello\\z")

	processor := &HelloWorldProcessor{}

	return &Bot{
		name:      "simplebot",
		in:        in,
		out:       out,
		checker:   checker,
		processor: processor,
	}
}

// "大吉", "吉", "中吉", "小吉", "末吉", "凶"のいずれかをランダムで返すbot
func NewOmikujiBot(out chan *model.Message) *Bot {
	in := make(chan *model.Message)

	checker := NewRegexpChecker("\\Aomikuji\\z")

	processor := &OmikujiProcessor{}

	return &Bot{
		name:      "omikujibot",
		in:        in,
		out:       out,
		checker:   checker,
		processor: processor,
	}
}

func (b *Bot) respond(m *model.Message) {
	message := b.processor.Process(m)
	b.out <- message
	fmt.Printf("%s send: %v\n", b.name, message)
}
