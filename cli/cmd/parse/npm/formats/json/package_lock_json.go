package json

import (
	"encoding/json"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	. "github.com/omniversion/omniversion/cli/types"
	"sort"
)

type PackageLockJson struct {
	Name                 string
	Version              string
	LockfileVersion      int
	Requires             bool
	Packages             map[string]Package
	License              string
	Dependencies         map[string]ResolvedDependency
	DevDependencies      map[string]ResolvedDependency
	PeerDependencies     map[string]ResolvedDependency
	OptionalDependencies map[string]ResolvedDependency
}

type Package struct {
	Name             string
	Version          string
	License          string
	Dependencies     map[string]string
	DevDependencies  map[string]string
	PeerDependencies map[string]string

	Resolved             string
	Integrity            string
	Dev                  bool
	OptionalDependencies map[string]string
	Peer                 bool
}

type ResolvedDependency struct {
	Bin                  map[string]string
	Dependencies         map[string]string
	Dev                  bool
	Engines              map[string]string
	Funding              map[string]string
	Integrity            string
	OptionalDependencies map[string]string
	Peer                 bool
	PeerDependencies     map[string]string
	PeerDependenciesMeta map[string]struct {
		Optional bool
	}
	Resolved string
	Version  string
}

func ParsePackageLockJsonFile(input string, _ stderr.Output) ([]PackageMetadata, error) {
	var result []PackageMetadata
	listJson := PackageLockJson{}
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &listJson)
	if jsonUnmarshallErr != nil {
		return result, jsonUnmarshallErr
	}

	if listJson.Version != "" && listJson.Name != "" {
		newItem := item.New(listJson.Name)
		newItem.Wanted = listJson.Version
		result = append(result, *newItem)
	}

	for name, resolvedDependency := range listJson.Dependencies {
		newItem := item.New(name)
		newItem.Locked = resolvedDependency.Version
		if resolvedDependency.Dev {
			newItem.Type = DevDependency
		} else if resolvedDependency.Peer {
			newItem.Type = PeerDependency
		} else {
			newItem.Type = ProdDependency
		}
		result = append(result, *newItem)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}
