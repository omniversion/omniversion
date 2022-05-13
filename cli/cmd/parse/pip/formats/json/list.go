package json

import (
	"encoding/json"
	"fmt"
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/item"
	. "github.com/omniversion/omniversion/cli/types"
	"sort"
)

type ListJsonOutput struct {
	Name                    string
	Version                 string
	EditableProjectLocation string `json:"editable_project_location"`
}

func ParseListOutput(input string) ([]PackageMetadata, error) {
	var result []PackageMetadata
	var listJson []ListJsonOutput
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &listJson)
	if jsonUnmarshallErr != nil {
		return result, fmt.Errorf("unable to parse JSON: %q", jsonUnmarshallErr)
	}

	for _, listItem := range listJson {
		newItem := item.New(listItem.Name)
		newItem.Current = listItem.Version
		newItem.Installations = []InstalledPackage{{
			Version:  listItem.Version,
			Location: listItem.EditableProjectLocation,
		}}
		result = append(result, *newItem)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}
