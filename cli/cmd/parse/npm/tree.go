package npm

import (
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
)

func isInTreeFormat(input string) bool {
	return regexp.MustCompile("(?m)^└──|^└─┬|^`--|^`-+").MatchString(input)
}

func parseAsTree(input string, result *[]PackageMetadata) *multierror.Error {
	treeLineRegex := regexp.MustCompile("(?m)^(?P<tree>[├|└─┬`\\- ]*) (?P<name>.*)@(?P<version>\\S*)(?P<deduped> deduped)?( (?P<addition>.*))?")
	items := treeLineRegex.FindAllStringSubmatch(input, -1)
	for _, foundItem := range items {
		// TODO: use indentation level to determine inter-package dependencies
		//treeContent := foundItem[treeLineRegex.SubexpIndex("tree")]
		name := foundItem[treeLineRegex.SubexpIndex("name")]
		version := foundItem[treeLineRegex.SubexpIndex("version")]
		newItem := PackageMetadata{
			Name:   name,
			Wanted: version,
		}
		if shared.InjectPackageManager {
			newItem.PackageManager = "npm"
		}
		*result = append(*result, newItem)
	}
	return nil
}
