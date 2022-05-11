package _default

import (
	"fmt"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
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
		return result, fmt.Errorf("invalid version data: %v", yamlUnmarshallErr)
	}
	for packageName, version := range npmVersionData {
		newResult := item.New(packageName)
		newResult.Current = version

		result = append(result, *newResult)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}
