package bot

import (
	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

const (
	postUrl = "http://localhost:8080/api/messages"
)

type (
	// Poster はInに渡されたmessageをPOSTするための構造体です
	Poster struct {
		In chan *model.Message
	}
)

// Run はPosterを起動する
func (p *Poster) Run() {
	for m := range p.In {
		out := &model.Message{}
		go postJson(postUrl, m, out)
	}
}

// NewPoster は新しいPoster構造体のポインタを返します
func NewPoster(bufferSize int) *Poster {
	in := make(chan *model.Message, bufferSize)
	return &Poster{
		In: in,
	}
}
