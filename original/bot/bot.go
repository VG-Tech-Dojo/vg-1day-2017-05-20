package bot

import (
	"context"
	"fmt"

	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

type Bot interface {
	Watch(context.Context)
	Respond()
	Run(context.Context)
}

type SimpleBot struct {
	name string
	in   chan *model.Message
	out  chan *model.Message
}

// Watchは投稿されたメッセージをチェックし続ける処理です
func (b *SimpleBot) Watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case m := <-b.in:
			fmt.Printf("%v", m)
			return
		}
	}
}

func (b *SimpleBot) Respond() {}

func NewSimpleBot(in chan *model.Message) *SimpleBot {
	return &SimpleBot{
		name: "simplebot",
		in:   in,
		out:  make(chan *model.Message),
	}
}

func (b *SimpleBot) Run(ctx context.Context) {
	fmt.Println("bot start")

	// メッセージ監視
	go b.Watch(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("bot stop")
			return
		}
	}
}
