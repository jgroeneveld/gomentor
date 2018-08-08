package quotes

import (
	"github.com/labstack/echo"
	"github.com/jgroeneveld/losmentor/fetching"
)

type Dependencies interface {
	JSONFetcher() fetching.JSONFetcher
}

type RandomQuoteController struct {
	Dependencies
}

func (c *RandomQuoteController) Handle(ctx echo.Context) error {
	quoteResponse, err := fetchRandomQuote(c.JSONFetcher())
	if err != nil {
		return err
	}

	quote := Quote{
		Author: quoteResponse.QuoteAuthor,
		Text:   quoteResponse.QuoteText,
	}

	return ctx.JSON(200, quote)
}
