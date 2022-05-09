package json

import (
	"encoding/json"
	"fmt"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	. "github.com/omniversion/omniversion/cli/types"
	"sort"
)

type OutdatedJsonOutput map[string]struct {
	Current   string
	Wanted    string
	Latest    string
	Dependent string
	Location  string
}

func ParseOutdatedOutput(input string, _ stderr.Output) ([]PackageMetadata, error) {
	var result []PackageMetadata
	outdatedJson := OutdatedJsonOutput{}
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &outdatedJson)
	if jsonUnmarshallErr != nil {
		return result, fmt.Errorf("unable to interpret this input! %v", jsonUnmarshallErr)
	}

	for name, details := range outdatedJson {
		newItem := item.New(name)
		newItem.Current = details.Current
		newItem.Wanted = details.Wanted
		newItem.Latest = details.Latest
		result = append(result, *newItem)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}
