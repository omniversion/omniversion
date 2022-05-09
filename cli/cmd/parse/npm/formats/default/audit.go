package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
)

func ParseAuditOutput(input string, _ stderr.Output) ([]PackageMetadata, error) {
	var result []PackageMetadata
	listRegex := regexp.MustCompile(`(?m)^(?P<name>\S+) {2}(?P<version>.*)\n\s*(Severity: (?P<severity>.*)|Depends on vulnerable versions of (?P<dependency>.*))`)
	items := listRegex.FindAllStringSubmatch(input, -1)
	for _, foundItem := range items {
		name := foundItem[listRegex.SubexpIndex("name")]
		version := foundItem[listRegex.SubexpIndex("version")]
		severity := foundItem[listRegex.SubexpIndex("severity")]
		newItem := item.New(name)
		newItem.Advisories = []Advisory{{
			Severity:           severity,
			VulnerableVersions: version,
		}}
		result = append(result, *newItem)
	}
	return result, nil
}
