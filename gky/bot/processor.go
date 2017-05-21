package bot

import (
	"regexp"
	"strings"

	"fmt"
	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/gky/env"
	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/gky/model"
  "net/url"
)

const (
	keywordApiUrlFormat = "https://jlp.yahooapis.jp/KeyphraseService/V1/extract?appid=%s&sentence=%s&output=json"
  chatBotApiUrl = "https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk"
)

type (
  TalkApiResponse struct {
    Status int `json:"status"`
    Message string `json:"message"`
    Results []struct {
      Perplexity float64 `json:"perplexity"`
      Reply string `json:"reply"`
    } `json:"results"`
  }
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
	KeywordProcessor struct{}

  UranaiProcessor struct{}
)

// Process は"hello, world!"というbodyがセットされたメッセージのポインタを返します
func (p *HelloWorldProcessor) Process(msgIn *model.Message) *model.Message {
	return &model.Message{
    SenderName: "bot",
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

func (p *UranaiProcessor) Process(msgIn *model.Message) *model.Message {
  gachas := []string{
    "SSレア",
    "Sレア",
    "レア",
    "ノーマル",
  }
  result := gachas[randIntn(len(gachas))]
  return &model.Message{
    SenderName: "bot",
    Body: result,
  }
}

func (p *ChatBotProcessor) Process(msgIn *model.Message) *model.Message {
	r := regexp.MustCompile("\\Atalk (.*)\\z")
	matchedStrings := r.FindStringSubmatch(msgIn.Body)
	text := matchedStrings[1]

  params := url.Values{
    "apikey": {env.ChatBotApiAppKey},
    "query": {text},
  }

	json := TalkApiResponse{}
	post(chatBotApiUrl, params, &json)

	return &model.Message{
    SenderName: "bot",
		Body: json.Results[0].Reply,
	}
}

