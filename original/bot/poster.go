package bot

import (
	"context"

	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
	"fmt"
)

const (
	postUrl = "http://localhost:8080/api/messages"
)

type (
	// Poster
	// Inputに入ったメッセージをPostする
	Poster struct {
		Input chan *model.Message
	}
)

func (p *Poster) Run(ctx context.Context) {
	for m := range p.Input {
		output := model.Message{}
		go postJson(postUrl, m, &output)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("poster stop")
			return
		}
	}
}

// posterの生成はこの関数を使う
func NewPoster(buffer_size int) *Poster {
	ch := make(chan *model.Message, buffer_size)
	return &Poster{
		Input: ch,
	}
}
