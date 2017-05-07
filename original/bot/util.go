package bot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"math/rand"
	"time"
	"github.com/pkg/errors"
	"fmt"

	"github.com/VG-Tech-Dojo/vg-1day-2017/original/env"
)

const (
	keywordApiUrlFormat = "https://jlp.yahooapis.jp/KeyphraseService/V1/extract?appid=%s&sentence=%s&output=json"
)

// inputをJSON形式でurlにPOSTする
func postJson(url string, input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return errors.Wrapf(err, "failed to decode json. data: %v", input)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return errors.Wrapf(err, "POST message: %v", data)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed to read response. response: %v", resp)
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return errors.Wrapf(err, "failed to encode json. json: %v", &output)
	}

	return nil
}

// 0からn-1までのintの乱数を返す
func randIntn (n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func extractKeyword(text string) []string {
	url := fmt.Sprintf(keywordApiUrlFormat, env.KeywordApiAppId, text)

	// TODO: エラー処理
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	respJson := map[string]int{}
	json.Unmarshal(respBody, &respJson)

	keywords := []string{}
	for keyword := range respJson {
		keywords = append(keywords, keyword)
	}

	return keywords
}
