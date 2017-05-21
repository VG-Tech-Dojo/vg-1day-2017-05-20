package bot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// get はurlにGETします
func get(url string, out interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return err
	}

	return nil
}

// post はurlにparamsをPOSTします
func post(url string, params url.Values, out interface{}) error {
	resp, err := http.PostForm(url, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return err
	}

	return nil
}

// postJSON はinputをJSON形式でurlにPOSTします
func postJSON(url string, input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	return nil
}

// randIntn は0からn-1までのintの乱数を返します
func randIntn(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
