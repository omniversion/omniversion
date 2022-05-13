package pip

import (
	"fmt"
	. "github.com/omniversion/omniversion/cli/cmd/parse/pip/formats"
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/formats/default"
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/formats/freeze"
	_json "github.com/omniversion/omniversion/cli/cmd/parse/pip/formats/json"
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
		}
	case JsonFormat:
		switch verb {
		case ListCommand:
			return _json.ParseListOutput
		}
	case FreezeFormat:
		switch verb {
		case ListCommand:
			return freeze.ParseListOutput
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
	Use:   "pip",
	Short: "Parse the output of `pip`",
	Long:  "Translate the output of `pip` into the omniversion format.",
	Run:   shared.WrapCommand(ParseOutput),
}
