package bot

import (
	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

const (
	postUrl = "http://localhost:8080/api/messages"
)

type (
	// Poster
	// Inに入ったメッセージをPOSTする
	Poster struct {
		In chan *model.Message
	}
)

func (p *Poster) Run() {
	for m := range p.In {
		out := &model.Message{}
		go postJson(postUrl, m, out)
	}
}

// posterの生成はこの関数を使う
func NewPoster(bufferSize int) *Poster {
	in := make(chan *model.Message, bufferSize)
	return &Poster{
		In: in,
	}
}
