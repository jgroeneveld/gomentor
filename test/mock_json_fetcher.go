package test

import (
	"encoding/json"
	"github.com/pkg/errors"
)

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
