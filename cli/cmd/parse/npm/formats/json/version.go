package json

import (
	_default "github.com/omniversion/omniversion/cli/cmd/parse/npm/formats/default"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	. "github.com/omniversion/omniversion/cli/types"
)

func ParseVersionOutput(input string, stderrOutput stderr.Output) ([]PackageMetadata, error) {
	// the default parser expects invalid JSON with unquoted keys, but can deal with this case, too
	return _default.ParseVersionOutput(input, stderrOutput)
}
