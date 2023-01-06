package main

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/pretty"
)

var (
	PrettyPrintJS = PrettyPrintJSJsonit
	jsonit        = jsoniter.ConfigCompatibleWithStandardLibrary
	prettyOpts    = pretty.Options{
		Width:    80,
		Prefix:   "",
		Indent:   "  ",
		SortKeys: true,
	}
)

func PrettyPrintJSJsonit(js []byte) []byte {
	if !jsonit.Valid(js) {
		return js
	}
	return pretty.PrettyOptions(js, &prettyOpts)
}
