package bot

import (
	"context"
	"fmt"

	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

type Bot interface {
	Watch(context.Context)
	Respond(*model.Message)
	Run(context.Context)
}

type SimpleBot struct {
	name      string
	in        chan *model.Message
	out       chan *model.Message
	checker   Checker
	processor Processor
}

// Watchは投稿されたメッセージをチェックし続ける処理です
func (b *SimpleBot) Watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case m := <-b.in:
			fmt.Printf("%s received: %v\n", b.name, m)

			if b.checker.Check(m) {
				b.Respond(m)
			}
		}
	}
}

func (b *SimpleBot) Respond(m *model.Message) {
	message := b.processor.Process(m)
	b.out <- message
	fmt.Printf("%s send: %v\n", b.name, message)
}

func NewSimpleBot(in chan *model.Message, out chan *model.Message) *SimpleBot {
	checker := NewRegexpChecker("\\Ahello\\z")

	processor := &HelloWorldProcessor{}

	return &SimpleBot{
		name: "simplebot",
		in:   in,
		out:  out,
		checker: checker,
		processor: processor,
	}
}

func NewOmikujiBot(in chan *model.Message, out chan *model.Message) *SimpleBot {
	checker := NewRegexpChecker("\\Aomikuji\\z")

	processor := &OmikujiProcessor{}

	return &SimpleBot{
		name: "omikujibot",
		in:   in,
		out:  out,
		checker: checker,
		processor: processor,
	}
}

func (b *SimpleBot) Run(ctx context.Context) {
	fmt.Printf("%s start\n", b.name)

	// メッセージ監視
	go b.Watch(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s stop", b.name)
			return
		}
	}
}
