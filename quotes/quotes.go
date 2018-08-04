package quotes

import (
	"github.com/labstack/echo"
)

func RandomQuote(ctx echo.Context) error {
	quoteResponse, err := fetchRandomQuote()
	if err != nil {
		return err
	}

	quote := Quote{
		Author: quoteResponse.QuoteAuthor,
		Text:   quoteResponse.QuoteText,
	}

	return ctx.JSON(200, quote)
}
