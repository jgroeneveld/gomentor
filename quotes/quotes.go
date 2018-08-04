package quotes

import (
	"github.com/labstack/echo"
)

func RandomQuote(ctx echo.Context) error {
	fetcher := NewJSONFetcher()

	quoteResponse, err := fetchRandomQuote(fetcher)
	if err != nil {
		return err
	}

	quote := Quote{
		Author: quoteResponse.QuoteAuthor,
		Text:   quoteResponse.QuoteText,
	}

	return ctx.JSON(200, quote)
}
