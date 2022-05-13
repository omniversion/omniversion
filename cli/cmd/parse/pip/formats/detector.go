package formats

import (
	"encoding/json"
	"regexp"
	"strings"
)

type InputFormat string
type Verb string

const (
	DefaultFormat InputFormat = "default"
	JsonFormat    InputFormat = "json"
	// FreezeFormat is also suitable for parsing `requirements.txt` files
	FreezeFormat InputFormat = "freeze"
)

const (
	ListCommand     Verb = "list"
	OutdatedCommand Verb = "outdated"
	VersionCommand  Verb = "version"
)

const VersionOutputRegex = `(?m)^pip (?P<version>\S+) from (?P<location>\S+) \(python (?P<pythonVersion>\S+)\)$`

func DetectVerbAndFormat(input string) (Verb, InputFormat) {
	if regexp.MustCompile(VersionOutputRegex).MatchString(input) {
		return VersionCommand, DefaultFormat
	}

	firstLine := strings.Split(input, "\n")[0]
	if strings.HasPrefix(firstLine, "Package ") {
		if strings.Contains(firstLine, "Latest") {
			return OutdatedCommand, DefaultFormat
		}
		return ListCommand, DefaultFormat
	}
	var jsonData []interface{}
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &jsonData)
	if jsonUnmarshallErr == nil {
		if strings.Contains(input, "\"latest_version\"") {
			return OutdatedCommand, JsonFormat
		}
		return ListCommand, JsonFormat
	}
	// `pip list --outdated --format=freeze` only shows the currently installed versions,
	// just like `pip list --format=freeze`, so we have no way of telling the two apart
	return ListCommand, FreezeFormat
}
