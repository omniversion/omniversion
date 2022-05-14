package formats

import (
	"encoding/json"
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
	var jsonData map[string]interface{}
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &jsonData)
	if jsonUnmarshallErr == nil {
		return ListCommand, JsonFormat
	}
	if strings.HasPrefix(input, "module ") {
		return GoModFile, DefaultFormat
	}
	return GoSumFile, DefaultFormat
}
