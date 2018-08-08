package quotes

import (
	"testing"

	"github.com/jgroeneveld/losmentor/test"
	"net/http/httptest"
	"github.com/labstack/echo"
	"net/http"
	"github.com/jgroeneveld/trial/assert"
	"github.com/jgroeneveld/schema"
)

func TestRandomQuoteController_Handle(t *testing.T) {
	fetcher := test.NewMockJSONFetcher()
	fetcher.Mock("https://api.forismatic.com/api/1.0/?method=getQuote&key=457635&format=json&lang=en", `
{
	"quoteText": "hallo",
	"quoteAuthor": "Peter"
}
`)

	dependencies := &test.DependencyConfiguration{
		JsonFetcher: fetcher,
	}

	c := &RandomQuoteController{dependencies}

	rec, ctx := newTestServer()

	err := c.Handle(ctx)
	assert.MustBeNil(t, err)

	test.AssertJSONSchema(t, schema.Map{
		"text":   "hallo",
		"author": "Peter",
	}, rec.Body)
}

func newTestServer() (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	return rec, ctx
}
