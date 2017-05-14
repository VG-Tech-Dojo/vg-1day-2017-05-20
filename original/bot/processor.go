package bot

import (
	"regexp"
	"strings"

	"fmt"
	"github.com/VG-Tech-Dojo/vg-1day-2017/original/env"
	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

const (
	keywordApiUrlFormat = "https://jlp.yahooapis.jp/KeyphraseService/V1/extract?appid=%s&sentence=%s&output=json"
)

type (
	// messageを受け取り、投稿用messageを作るインターフェース
	Processor interface {
		Process(message *model.Message) *model.Message
	}

	// "hello, world!"メッセージを作るprocessor
	HelloWorldProcessor struct{}

	// "大吉", "吉", "中吉", "小吉", "末吉", "凶"のいずれかをランダムで作るprocessor
	OmikujiProcessor struct{}

	// メッセージ本文からキーワードを抽出するprocessor
	KeywordProcessor struct{}
)

// "hello, world!"メッセージを作る
func (p *HelloWorldProcessor) Process(msgIn *model.Message) *model.Message {
	return &model.Message{
		Body: msgIn.Body + ", world!",
	}
}

// "大吉", "吉", "中吉", "小吉", "末吉", "凶"のいずれかをランダムで作る
func (p *OmikujiProcessor) Process(msgIn *model.Message) *model.Message {
	fortunes := []string{
		"大吉",
		"吉",
		"中吉",
		"小吉",
		"末吉",
		"凶",
	}
	result := fortunes[randIntn(len(fortunes))]
	return &model.Message{
		Body: result,
	}
}

// メッセージ本文からキーワードを抽出する
func (p *KeywordProcessor) Process(msgIn *model.Message) *model.Message {
	r := regexp.MustCompile("\\Akeyword (.*)\\z")
	matchedStrings := r.FindStringSubmatch(msgIn.Body)
	text := matchedStrings[1]

	url := fmt.Sprintf(keywordApiUrlFormat, env.KeywordApiAppId, text)

	json := map[string]int{}
	getJSON(url, &json)

	keywords := []string{}
	for keyword := range map[string]int(json) {
		keywords = append(keywords, keyword)
	}

	return &model.Message{
		Body: "キーワード：" + strings.Join(keywords, ", "),
	}
}
