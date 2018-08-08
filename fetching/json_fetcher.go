package fetching

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type JSONFetcher interface {
	Get(url string, out interface{}) error
}

type JSONFetcherImpl struct {
	Client *http.Client
}

func NewJSONFetcher() *JSONFetcherImpl {
	return &JSONFetcherImpl{Client: &http.Client{}}
}

func (fetcher *JSONFetcherImpl) Get(url string, out interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := fetcher.Client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	return nil
}