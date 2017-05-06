package bot

import (
	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
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

func (p *Poster) Run() {
	for m := range p.Input {
		output := model.Message{}
		go postJson(postUrl, m, &output)
	}
}

// posterの生成はこの関数を使う
func NewPoster(buffer_size int) *Poster {
	ch := make(chan *model.Message, buffer_size)
	return &Poster{
		Input: ch,
	}
}
