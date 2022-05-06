package npm

import (
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func isInOutdatedTableFormat(input string) bool {
	firstLine := strings.Split(input, "\n")[0]
	return strings.Contains(firstLine, "Package") &&
		strings.Contains(firstLine, "Current") &&
		strings.Contains(firstLine, "Wanted") &&
		strings.Contains(firstLine, "Latest")
}

func parseAsOutdatedTable(input string, result *[]PackageMetadata) *multierror.Error {
	listRegex := regexp.MustCompile(`(?m)(?P<name>\S+)\s+(?P<current>\S+)\s+(?P<wanted>\S+)\s+(?P<latest>\S+)\s+(?P<location>\S+)\s+(?P<dependedBy>.*)$`)
	items := listRegex.FindAllStringSubmatch(input, -1)
	var allErrors *multierror.Error = nil
	for index, foundItem := range items {
		if index == 0 {
			// ignore the first line consisting of headers
			continue
		}
		name := foundItem[listRegex.SubexpIndex("name")]
		current := foundItem[listRegex.SubexpIndex("current")]
		wanted := foundItem[listRegex.SubexpIndex("wanted")]
		latest := foundItem[listRegex.SubexpIndex("latest")]
		location := foundItem[listRegex.SubexpIndex("location")]
		newItem := PackageMetadata{
			Name:    name,
			Current: current,
			Wanted:  wanted,
			Latest:  latest,
		}
		if current != "" {
			newItem.Installations = []InstalledPackage{{
				Location: location,
				Version:  current}}
		}
		if shared.InjectPackageManager {
			newItem.PackageManager = "npm"
		}
		*result = append(*result, newItem)
	}
	return allErrors
}
