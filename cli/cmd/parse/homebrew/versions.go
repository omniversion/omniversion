package homebrew

import (
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
)

func parseAsListWithVersionsCommandOutput(input string) ([]PackageMetadata, error) {
	extractionRegex := regexp.MustCompile(`(?m)(?P<name>\S*) (?P<version>.*)?`)
	items := extractionRegex.FindAllStringSubmatch(input, -1)

	result := make([]PackageMetadata, 0, len(items))
	var allErrors *multierror.Error
	for _, item := range items {
		name := item[extractionRegex.SubexpIndex("name")]
		version := item[extractionRegex.SubexpIndex("version")]
		newItem := PackageMetadata{
			Name:          name,
			Current:       version,
			Installations: []InstalledPackage{{Version: version}},
		}
		if shared.InjectPackageManager {
			newItem.PackageManager = "rubygems"
		}
		result = append(result, newItem)
	}
	return result, allErrors.ErrorOrNil()
}
