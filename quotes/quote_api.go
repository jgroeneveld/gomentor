package quotes

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type quoteApiResponse struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
}

func fetchRandomQuote() (*quoteApiResponse, error) {
	url := "https://api.forismatic.com/api/1.0/?method=getQuote&key=457635&format=json&lang=en"

	quoteResponse := &quoteApiResponse{}

	err := getJSON(url, quoteResponse)
	if err != nil {
		return nil, err
	}

	return quoteResponse, nil
}

func getJSON(url string, target interface{}) error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}
