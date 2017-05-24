package bot

import (
	"regexp"
	"strings"

	"fmt"
	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/team1/env"
	"github.com/VG-Tech-Dojo/vg-1day-2017-05-20/team1/model"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const (
	keywordApiUrlFormat = "https://jlp.yahooapis.jp/KeyphraseService/V1/extract?appid=%s&sentence=%s&output=json"
	yahooAuctionApiUrlFormat = "https://auctions.yahooapis.jp/AuctionWebService/V2/json/search?appid=%s&query=%s"
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

	YahooAuctionProcessor struct{}
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

func (p *YahooAuctionProcessor) Process(msgIn *model.Message) *model.Message {
	r := regexp.MustCompile("\\Afollow (.*)\\z")
	matchedStrings := r.FindStringSubmatch(msgIn.Body)
	text := matchedStrings[1]

	url := fmt.Sprintf(yahooAuctionApiUrlFormat, env.YahooAuctionApiAppId, text)

	type (
		Item struct {
			AuctionID string `json:"AuctionID"`
			Title string `json:"title"`
			Seller map[string]string `json:"Seller"`
			ItemUrl string `json:"ItemUrl"`
			AuctionItemUrl string `json:"AuctionItemUrl"`
			Image string `json:"Image"`
			OriginalImageNum string `json:"OriginalImageNum"`
			CurrentPrice string `json:"CurrentPrice"`
			Bids string `json:"Bids"`
			EndTime string `json:"EndTime"`
			IsReserved string `json:"IsReserved"`
			CharityOption map[string] string `json:"CharityOption"`
			Option map[string]string `json:"Option"`
			IsAdult string `json:"IsAdult"`
		}

		Result struct {
			UnitsWord []string `json:"UnitsWord"`
			Item []Item `json:"Item"`
		}

		ResultSet struct {
			Attributes map[string]string `json:"@attributes"`
			Result Result `json:"Result"`
		}

		Response struct {
			ResultSet ResultSet `json:"ResultSet"`
		}
	)

	resp := Response{}
	getYahooAuctionAPI(url, &resp)

	return &model.Message{
		Body: "[" + resp.ResultSet.Result.Item[0].Title + "] " + resp.ResultSet.Result.Item[0].AuctionItemUrl,
		Image: resp.ResultSet.Result.Item[0].Image,
	}
}

func getYahooAuctionAPI(url string, out interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	stringResp := strings.TrimRight(strings.TrimLeft(string(respBody), "loaded("), ")")
	respBody = []byte(stringResp)

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return err
	}

	return nil
}
