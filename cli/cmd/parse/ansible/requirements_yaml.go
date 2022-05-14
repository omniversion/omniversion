package ansible

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/ansible/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"gopkg.in/yaml.v3"
)

type RequirementYamlItem struct {
	Include string
	Name    string
	Scm     string
	Src     string
	Version string
}

func ParseRequirementsYamlFile(input string) ([]PackageMetadata, error) {
	var parsedFile []RequirementYamlItem
	err := yaml.Unmarshal([]byte(input), &parsedFile)
	if err != nil {
		return []PackageMetadata{}, err
	}
	var result []PackageMetadata
	for _, dataItem := range parsedFile {
		newItem := item.New(dataItem.Name)
		newItem.Wanted = shared.CleanVersion(dataItem.Version)
		newItem.Name = dataItem.Src
		// TODO: also parse files referenced in `data.Include`
		result = append(result, *newItem)
	}
	return result, nil
}
