package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
)

func ParseOutdatedOutput(input string, _ stderr.Output) ([]PackageMetadata, error) {
	var result []PackageMetadata
	listRegex := regexp.MustCompile(`(?m)(?P<name>\S+)\s+(?P<current>\S+)\s+(?P<wanted>\S+)\s+(?P<latest>\S+)\s+(?P<location>\S+)\s+(?P<dependedBy>.*)$`)
	items := listRegex.FindAllStringSubmatch(input, -1)
	for index, foundItem := range items {
		if index == 0 {
			// ignore the first line consisting of headers
			continue
		}
		name := foundItem[listRegex.SubexpIndex("name")]
		current := foundItem[listRegex.SubexpIndex("current")]
		if current == "MISSING" {
			current = ""
		}
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
		result = append(result, newItem)
	}
	return result, nil
}
