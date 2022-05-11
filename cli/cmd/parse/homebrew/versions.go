package homebrew

import (
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/homebrew/item"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
)

func parseAsListWithVersionsCommandOutput(input string) ([]PackageMetadata, error) {
	extractionRegex := regexp.MustCompile(`(?m)(?P<name>\S*) (?P<version>.*)?`)
	items := extractionRegex.FindAllStringSubmatch(input, -1)

	result := make([]PackageMetadata, 0, len(items))
	var allErrors *multierror.Error
	for _, itemData := range items {
		name := itemData[extractionRegex.SubexpIndex("name")]
		version := itemData[extractionRegex.SubexpIndex("version")]
		newItem := item.New(name)
		newItem.Current = version
		newItem.Installations = []InstalledPackage{{Version: version}}

		result = append(result, *newItem)
	}
	return result, allErrors.ErrorOrNil()
}
