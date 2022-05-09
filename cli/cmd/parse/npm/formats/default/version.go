package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"gopkg.in/yaml.v3"
	"sort"
)

type VersionJson map[string]string

func ParseVersionOutput(input string, _ stderr.Output) ([]PackageMetadata, error) {
	var result []PackageMetadata
	// the keys in the default `npm version` output are unquoted,
	// but the yaml parser can deal with this
	npmVersionData := VersionJson{}
	yamlUnmarshallErr := yaml.Unmarshal([]byte(input), &npmVersionData)
	if yamlUnmarshallErr != nil {
		return result, yamlUnmarshallErr
	}
	for packageName, version := range npmVersionData {
		newResult := PackageMetadata{
			Name:    packageName,
			Current: version,
		}
		if shared.InjectPackageManager {
			newResult.PackageManager = "npm"
		}
		result = append(result, newResult)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}
