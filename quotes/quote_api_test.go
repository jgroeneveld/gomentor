package quotes

import (
		"testing"
		"github.com/jgroeneveld/trial/assert"
		"github.com/jgroeneveld/losmentor/test"
)

func Test_fetchRandomQuote(t *testing.T) {
	fetcher := test.NewMockJSONFetcher()

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


