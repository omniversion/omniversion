package json

import (
	"encoding/json"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	. "github.com/omniversion/omniversion/cli/types"
	"sort"
)

type ListJsonOutput struct {
	Dependencies map[string]Dependency
	From         string
	Name         string
	PeerMissing  bool
	Problems     []string
	Required     Requirement
	Resolved     string
	Version      string
}

func ParseListOutput(input string, stderrOutput stderr.Output) ([]PackageMetadata, error) {
	var result []PackageMetadata
	listJson := ListJsonOutput{}
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &listJson)
	if jsonUnmarshallErr != nil {
		return result, jsonUnmarshallErr
	}

	if listJson.Name != "" {
		newItem := item.New(listJson.Name)
		newItem.Current = listJson.Version
		result = append(result, *newItem)
	}

	for name, dependency := range listJson.Dependencies {
		newItem := item.New(name)
		newItem.Wanted = dependency.Version
		if dependency.Version == "" {
			if dependency.Extraneous {
				newItem.Extraneous = true
			} else {
				// missing (i.e. wanted, but not installed) packages don't always report a version,
				// but we can get it from the stderr output
				for _, stderrError := range stderrOutput.Errors {
					if stderrError.Problem == "missing" && stderrError.Name == name {
						newItem.Wanted = stderrError.Version
						newItem.Missing = true
					}
				}
			}
		} else {
			newItem.Current = dependency.Version
		}
		if len(dependency.Dependencies) > 0 {
			newItem.Dependencies = make([]string, 0, len(dependency.Dependencies))
			for subDependencyName := range dependency.Dependencies {
				if subDependencyName != "" {
					newItem.Dependencies = append(newItem.Dependencies, subDependencyName)
				}
			}
		}
		result = append(result, *newItem)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}
