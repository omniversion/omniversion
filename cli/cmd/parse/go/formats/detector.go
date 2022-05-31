package formats

import (
	"strings"
)

type InputFormat string
type Verb string

const (
	DefaultFormat InputFormat = "default"
	JsonFormat    InputFormat = "json"
)

const (
	ListCommand    Verb = "list"
	VersionCommand Verb = "version"
	GoModFile      Verb = "go.mod"
	GoSumFile      Verb = "go.sum"
)

func DetectVerbAndFormat(input string) (Verb, InputFormat) {
	if strings.HasPrefix(input, "go version ") {
		return VersionCommand, DefaultFormat
	}
	if strings.HasPrefix(input, "{") {
		return ListCommand, JsonFormat
	}
	if strings.HasPrefix(input, "module ") {
		return GoModFile, DefaultFormat
	}
	if strings.Contains(input, " ") {
		return GoSumFile, DefaultFormat
	}
	return ListCommand, DefaultFormat
}
