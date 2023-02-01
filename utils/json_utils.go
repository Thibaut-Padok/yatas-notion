package utils

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/pretty"
)

var (
	PrettyPrintJS = PrettyPrintJSJsonit
	Jsonit        = jsoniter.ConfigCompatibleWithStandardLibrary
	prettyOpts    = pretty.Options{
		Width:    80,
		Prefix:   "",
		Indent:   "  ",
		SortKeys: true,
	}
)

func PrettyPrintJSJsonit(js []byte) []byte {
	if !Jsonit.Valid(js) {
		return js
	}
	return pretty.PrettyOptions(js, &prettyOpts)
}
