package bot

import (
	"regexp"
	"strings"

	"fmt"
	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/team02/env"
	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/team02/model"
)

const (
	keywordApiUrlFormat = "https://jlp.yahooapis.jp/KeyphraseService/V1/extract?appid=%s&sentence=%s&output=json"
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

	HomeProcessor struct {}

	HomeJson struct {
		Humidity float64 `json:"humidity"`
		Temperature float64 `json:"temperature"`
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

func (p *HomeProcessor) Process(msgIn *model.Message) *model.Message {

	json := HomeJson{}
	get("http://192.168.100.150:5000", &json)

	temp := json.Temperature
	humid := json.Humidity

	DI := 0.81 * temp + 0.01 * humid * (0.99 * temp - 14.3) + 46.3
	var msg string
	if DI < 55 {
		msg = "寒い"
	} else if DI < 60 {
		msg = "肌寒い"
	} else if DI < 65 {
		msg = "何も感じない"
	} else if DI < 70{
		msg = "快い"
	} else if DI < 75{
		msg = "暑くない"
	} else if DI < 80{
		msg = "やや暑い"
	} else if DI < 85{
		msg = "暑くて汗が出る"
	} else {
		msg = "暑くてたまらない"
	}
	msg = fmt.Sprintf("温度: %.0f°, 湿度: %.0f%%, 不快指数: %.3f, メッセージ: %s", temp, humid, DI, msg)
	return &model.Message{
		Body: "home：" + msg,
	}
}
