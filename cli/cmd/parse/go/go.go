package _go

import (
	"fmt"
	. "github.com/omniversion/omniversion/cli/cmd/parse/go/formats"
	_default "github.com/omniversion/omniversion/cli/cmd/parse/go/formats/default"
	_json "github.com/omniversion/omniversion/cli/cmd/parse/go/formats/json"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
)

func parserForVerbAndFormat(verb Verb, format InputFormat) func(string) ([]PackageMetadata, error) {
	switch format {
	case DefaultFormat:
		switch verb {
		case ListCommand:
			return _default.ParseListOutput
		case VersionCommand:
			return _default.ParseVersionOutput
		case GoModFile:
			return _default.ParseGoModFile
		case GoSumFile:
			return _default.ParseGoSumFile
		}
	case JsonFormat:
		switch verb {
		case ListCommand:
			return _json.ParseListOutput
		}
	}
	return func(string) ([]PackageMetadata, error) {
		return []PackageMetadata{}, fmt.Errorf("unrecognized verb/format combination: %q / %q", verb, format)
	}
}

func ParseOutput(input string) ([]PackageMetadata, error) {
	verb, format := DetectVerbAndFormat(input)
	result, parseErr := parserForVerbAndFormat(verb, format)(input)

	return result, parseErr
}

var ParseCommand = &cobra.Command{
	Use:     "go",
	Short:   "Parse the output of go CLI commands and/or `go.mod`/`go.sum` files",
	Long:    `Translate the output of npm into the omniversion format.`,
	Run:     shared.WrapCommand(ParseOutput),
	Aliases: []string{"go", "mod"},
}
