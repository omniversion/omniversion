package brew

import (
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/brew/item"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"sort"
	"strings"
)

func isDefaultListOutput(input string) bool {
	return strings.HasPrefix(input, "==> Formulae") || strings.HasPrefix(input, "\u001B[34m==>")
}

func parseAsDefaultListOutput(input string) ([]PackageMetadata, error) {
	extractionRegex := regexp.MustCompile(`(?m)^(\[34m)?==>.*|(?P<name>[^= \n]\S+)\s+`)
	items := extractionRegex.FindAllStringSubmatch(input, -1)

	result := make([]PackageMetadata, 0, len(items))
	var allErrors *multierror.Error
	for _, itemData := range items {
		name := itemData[extractionRegex.SubexpIndex("name")]
		if name == "" {
			continue
		}
		newItem := item.New(name)
		newItem.Installations = []InstalledPackage{{}}
		result = append(result, *newItem)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, allErrors.ErrorOrNil()
}
