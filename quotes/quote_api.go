package quotes

import "github.com/jgroeneveld/losmentor/fetching"

type quoteApiResponse struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
}

func fetchRandomQuote(fetcher fetching.JSONFetcher) (*quoteApiResponse, error) {
	url := "https://api.forismatic.com/api/1.0/?method=getQuote&key=457635&format=json&lang=en"

	quoteResponse := &quoteApiResponse{}

	err := fetcher.Get(url, quoteResponse)
	if err != nil {
		return nil, err
	}

	return quoteResponse, nil
}


