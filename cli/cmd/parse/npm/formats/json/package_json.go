package json

import (
	"encoding/json"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	. "github.com/omniversion/omniversion/cli/types"
	"sort"
)

type PackageJson struct {
	Dependencies             map[string]string
	DevDependencies          map[string]string
	License                  string
	Name                     string
	OptionalDependencies     map[string]string
	OptionalPeerDependencies map[string]string
	PeerDependencies         map[string]string
	Scripts                  map[string]string
	Version                  string
}

func ParsePackageJsonFile(input string, _ stderr.Output) ([]PackageMetadata, error) {
	var result []PackageMetadata
	listJson := PackageJson{}
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &listJson)
	if jsonUnmarshallErr != nil {
		return result, jsonUnmarshallErr
	}

	if listJson.Version != "" && listJson.Name != "" {
		newItem := item.New(listJson.Name)
		newItem.Current = listJson.Version
		result = append(result, *newItem)
	}

	for name, version := range listJson.Dependencies {
		newItem := item.New(name)
		newItem.Wanted = version
		newItem.Type = ProdDependency
		result = append(result, *newItem)
	}
	for name, version := range listJson.DevDependencies {
		newItem := item.New(name)
		newItem.Wanted = version
		newItem.Type = DevDependency
		result = append(result, *newItem)
	}
	for name, version := range listJson.PeerDependencies {
		newItem := item.New(name)
		newItem.Wanted = version
		newItem.Type = PeerDependency
		result = append(result, *newItem)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}
