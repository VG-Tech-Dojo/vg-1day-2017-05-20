package bot

import (
	"context"
	"fmt"

	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

type Bot struct {
	name      string
	in        chan *model.Message
	out       chan *model.Message
	checker   Checker
	processor Processor
}

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
				b.Respond(m)
			}
		}
	}
}

func (b *Bot) Respond(m *model.Message) {
	message := b.processor.Process(m)
	b.out <- message
	fmt.Printf("%s send: %v\n", b.name, message)
}

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
