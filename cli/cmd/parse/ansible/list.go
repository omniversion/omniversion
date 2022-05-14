package ansible

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/ansible/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func ParseListOutput(input string) ([]PackageMetadata, error) {
	compiledRegex := regexp.MustCompile(`(?m)^- (?P<name>\S*), v?(?P<version>\S+)$`)
	matches := compiledRegex.FindAllStringSubmatch(input, -1)

	result := make([]PackageMetadata, 0, len(matches))
	for _, match := range matches {
		name := match[compiledRegex.SubexpIndex("name")]
		version := match[compiledRegex.SubexpIndex("version")]
		newItem := item.New(name)
		nameComponents := strings.Split(name, ".")
		shortName := nameComponents[len(nameComponents)-1]
		newItem.Aliases = []string{shortName}
		newItem.Current = shared.CleanVersion(version)
		newItem.Installations = []InstalledPackage{{
			Version: shared.CleanVersion(version),
		}}
		result = append(result, *newItem)
	}
	return result, nil
}
