package json

import (
	"encoding/json"
	"fmt"
	"github.com/omniversion/omniversion/cli/cmd/parse/pip/item"
	. "github.com/omniversion/omniversion/cli/types"
	"sort"
)

type OutdatedJsonOutput struct {
	Name                    string
	Version                 string
	EditableProjectLocation string `json:"editable_project_location"`
	LatestVersion           string `json:"latest_version"`
	LatestFiletype          string `json:"latest_filetype"`
}

func ParseOutdatedOutput(input string) ([]PackageMetadata, error) {
	var result []PackageMetadata
	var outdatedJson []OutdatedJsonOutput
	jsonUnmarshallErr := json.Unmarshal([]byte(input), &outdatedJson)
	if jsonUnmarshallErr != nil {
		return result, fmt.Errorf("unable to parse JSON: %q", jsonUnmarshallErr)
	}

	for _, outdatedItem := range outdatedJson {
		newItem := item.New(outdatedItem.Name)
		newItem.Current = outdatedItem.Version
		newItem.Latest = outdatedItem.LatestVersion
		newItem.Installations = []InstalledPackage{{
			Version:  outdatedItem.Version,
			Location: outdatedItem.EditableProjectLocation,
		}}
		result = append(result, *newItem)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}
