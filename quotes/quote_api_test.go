package quotes

import (
		"testing"
	"github.com/pkg/errors"
	"github.com/jgroeneveld/trial/assert"
	"encoding/json"
)

func Test_fetchRandomQuote(t *testing.T) {
	fetcher := NewMockJSONFetcher()

	fetcher.Mock("https://api.forismatic.com/api/1.0/?method=getQuote&key=457635&format=json&lang=en", `
{
	"quoteText": "hallo",
	"quoteAuthor": "Peter"
}
`)

	response, err := fetchRandomQuote(fetcher)
	assert.MustBeNil(t, err)

	assert.Equal(t, "Peter", response.QuoteAuthor)
	assert.Equal(t, "hallo", response.QuoteText)
}

type MockJSONFetcher struct {
	mocks map[string]string
}

func NewMockJSONFetcher() *MockJSONFetcher {
	return &MockJSONFetcher{
		mocks: map[string]string{},
	}
}

func (fetcher *MockJSONFetcher) Mock(url string, mock string) {
	fetcher.mocks[url] = mock
}

func (fetcher *MockJSONFetcher) Get(url string, out interface{}) error {
	if mock := fetcher.mocks[url]; mock != "" {
		err := json.Unmarshal([]byte(mock), out)
		if err != nil {
			panic(err.Error())
		}

		return nil
	}

	return errors.Errorf("No Mock for %s", url)
}
