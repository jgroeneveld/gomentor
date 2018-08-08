package test

import "github.com/jgroeneveld/losmentor/fetching"

type DependencyConfiguration struct {
	JsonFetcher fetching.JSONFetcher
}

func (deps *DependencyConfiguration) JSONFetcher() fetching.JSONFetcher {
	if deps.JsonFetcher == nil {
		panic("JSONFetcher not configured")
	}
	return deps.JsonFetcher
}
