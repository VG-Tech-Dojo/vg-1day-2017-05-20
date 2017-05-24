package bot

import (
	"encoding/xml"
	"regexp"
	"strings"

	"fmt"
	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/team3/env"
	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/team3/model"
)

const (
	keywordApiUrlFormat = "http://jlp.yahooapis.jp/MAService/V1/parse?appid=dj0zaiZpPVJ6YzZMeTVTbkJLZSZzPWNvbnN1bWVyc2VjcmV0Jng9MmE-&sentence="
)

type (
	// Processor はmessageを受け取り、投稿用messageを作るインターフェースです
	Processor interface {
		Process(message *model.Message) *model.Message
	}

	// HelloWorldProcessor は"hello, world!"メッセージを作るprocessorの構造体です
	HelloWorldProcessor struct{}

	// OmikujiProcessor は"大吉", "吉", "中吉", "小吉", "末吉", "凶"のいずれかをランダムで作るprocessorの構造体です
	OmikujiProcessor struct{}

	// メッセージ本文からキーワードを抽出するprocessorの構造体です
	KeywordProcessor     struct{}
	WordEmotionProcessor struct{}

	XML struct {
		surface []string `xml:"surface"`
		reading []string `xml:"reading"`
		pos     []string `xml:"pos"`
	}
)

// Process は"hello, world!"というbodyがセットされたメッセージのポインタを返します
func (p *HelloWorldProcessor) Process(msgIn *model.Message) *model.Message {
	return &model.Message{
		Body: msgIn.Body + ", world!",
	}
}

// Process は"大吉", "吉", "中吉", "小吉", "末吉", "凶"のいずれかがbodyにセットされたメッセージへのポインタを返します
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

// Process はメッセージ本文からキーワードを抽出します
func (p *KeywordProcessor) Process(msgIn *model.Message) *model.Message {
	r := regexp.MustCompile("\\Akeyword (.*)\\z")
	matchedStrings := r.FindStringSubmatch(msgIn.Body)
	text := matchedStrings[1]

	url := fmt.Sprintf(keywordApiUrlFormat, env.KeywordApiAppId, text)

	json := map[string]int{}
	get(url, &json)

	keywords := []string{}
	for keyword := range map[string]int(json) {
		keywords = append(keywords, keyword)
	}

	return &model.Message{
		Body: "キーワード：" + strings.Join(keywords, ", "),
	}
}

// Process はメッセージ本文からキーワードを抽出します
func (p *WordEmotionProcessor) Process(msgIn *model.Message) *model.Message {
	r := regexp.MustCompile("\\Atalk (.*)\\z")
	matchedStrings := r.FindStringSubmatch(msgIn.Body)
	text := matchedStrings[1]

	url := keywordApiUrlFormat + text

	result := XML{}
	xmlGet(url, &result)
	data := new(XML)
	if err := xml.Unmarshal(result, data); err != nil {
		fmt.Println("XML Unmarshal error", err)
		// return
	}
	// var result_text string
	//
	// for _, xml := range result {
	// 	fmt.Println(xml.surface)
	// 	result_text = xml.surface
	// }

	// keywords := []string{}
	// for keyword := range map[string]int(json) {
	// 	keywords = append(keywords, keyword)
	// }
	return &model.Message{
		Body: "キーワード： " + data.surface[0],
	}

}
