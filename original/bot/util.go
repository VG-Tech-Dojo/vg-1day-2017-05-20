package bot

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
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
		return errors.Wrapf(err, "GET url: %v", url)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed to read response. response: %v", resp)
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return errors.Wrapf(err, "failed to encode json. json: %v", &out)
	}

	return nil
}

// post はurlにparamsをPOSTします
func post(url string, params url.Values, out interface{}) error {
	resp, err := http.PostForm(url, params)
	if err != nil {
		return errors.Wrapf(err, "POST url: %v, params: %v", url, params)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed to read response. response: %v", resp)
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return errors.Wrapf(err, "failed to encode json. json: %v", &out)
	}

	return nil
}

// postJSON はinputをJSON形式でurlにPOSTします
func postJSON(url string, input interface{}, output interface{}) error {
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

// randIntn は0からn-1までのintの乱数を返します
func randIntn(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
