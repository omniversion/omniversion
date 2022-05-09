package homebrew

import (
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func isDefaultListOutput(input string) bool {
	return strings.HasPrefix(input, "==> Formulae")
}

func parseAsDefaultListOutput(input string) ([]PackageMetadata, error) {
	extractionRegex := regexp.MustCompile(`(?m)(?P<name>[^= \n]\S+)\s{2,}`)
	items := extractionRegex.FindAllStringSubmatch(input, -1)

	result := make([]PackageMetadata, 0, len(items))
	var allErrors *multierror.Error
	for _, item := range items {
		name := item[extractionRegex.SubexpIndex("name")]
		newItem := PackageMetadata{
			Name:          name,
			Installations: []InstalledPackage{{}},
		}
		if shared.InjectPackageManager {
			newItem.PackageManager = "rubygems"
		}
		result = append(result, newItem)
	}
	return result, allErrors.ErrorOrNil()
}