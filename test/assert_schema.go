package test

import (
	"testing"
	"github.com/jgroeneveld/schema"
	"io"
	"github.com/jgroeneveld/trial/th"
)

func AssertJSONSchema(t *testing.T, matcher schema.Matcher, r io.Reader) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		th.Error(t, 1, err.Error())
	}
}

func MustMatchJSONSchema(t *testing.T, matcher schema.Matcher, r io.Reader) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		th.Error(t, 1, err.Error())
		t.FailNow()
	}
}
