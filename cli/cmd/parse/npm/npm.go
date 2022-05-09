package npm

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	. "github.com/omniversion/omniversion/cli/cmd/parse/npm/formats"
	_default "github.com/omniversion/omniversion/cli/cmd/parse/npm/formats/default"
	_json "github.com/omniversion/omniversion/cli/cmd/parse/npm/formats/json"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/formats/parseable"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
)

func parserForVerbAndFormat(verb Verb, format InputFormat) func(string, stderr.Output) ([]PackageMetadata, error) {
	switch format {
	case DefaultFormat:
		switch verb {
		case AuditCommand:
			return _default.ParseAuditOutput
		case ListCommand:
			return _default.ParseListOutput
		case OutdatedCommand:
			return _default.ParseOutdatedOutput
		case VersionCommand:
			return _default.ParseVersionOutput
		}
	case JsonFormat:
		switch verb {
		case AuditCommand:
			return _json.ParseAuditOutput
		case ListCommand:
			return _json.ParseListOutput
		case OutdatedCommand:
			return _json.ParseOutdatedOutput
		case VersionCommand:
			return _json.ParseVersionOutput
		case PackageJsonFile:
			return _json.ParsePackageJsonFile
		case PackageLockJsonFile:
			return _json.ParsePackageLockJsonFile
		}
	case ParseableFormat:
		switch verb {
		// only `outdated` and `list` have parseable versions
		case ListCommand:
			return parseable.ParseListOutput
		case OutdatedCommand:
			return parseable.ParseOutdatedOutput
		}
	}
	return func(string, stderr.Output) ([]PackageMetadata, error) {
		return []PackageMetadata{}, fmt.Errorf("unrecognized verb/format combination: %q / %q", verb, format)
	}
}

func ParseOutput(rawInput string) ([]PackageMetadata, error) {
	// remove problems that might appear in stderr
	// and would prevent us from parsing the content as JSON
	// this is relevant if stdout and stderr have been merged,
	// e.g. in terminal output copied from the console
	input, stderrOutput, stripErr := stderr.Strip(rawInput)

	verb, format := formats.DetectVerbAndFormat(input, stderrOutput)
	result, parseErr := parserForVerbAndFormat(verb, format)(input, stderrOutput)

	return result, multierror.Append(stripErr, parseErr).ErrorOrNil()
}

var ParseCommand = &cobra.Command{
	Use:   "npm",
	Short: "Parse the output of npm",
	Long:  `Translate the output of npm into the omniversion format.`,
	Run:   shared.WrapCommand(ParseOutput),
}
