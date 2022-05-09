package parseable

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	. "github.com/omniversion/omniversion/cli/types"
	"sort"
	"strings"
)

func ParseListOutput(input string, stderrOutput stderr.Output) ([]PackageMetadata, error) {
	var result []PackageMetadata
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		components := strings.Split(line, "/")
		name := components[len(components)-1]
		newItem := item.New(name)
		newItem.Installations = []InstalledPackage{{
			Location: line,
		}}

		result = append(result, *newItem)
	}

	for _, stderrError := range stderrOutput.Errors {
		if stderrError.Problem == "missing" {
			newItem := item.New(stderrError.Name)
			newItem.Missing = true
			newItem.Wanted = stderrError.Version
			result = append(result, *newItem)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}
