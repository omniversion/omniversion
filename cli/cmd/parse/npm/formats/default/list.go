package _default

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm/stderr"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"unicode/utf8"
)

func ParseListOutput(input string, _ stderr.Output) ([]PackageMetadata, error) {
	var result []*PackageMetadata
	treeLineRegex := regexp.MustCompile("(?m)^(?P<treePrefix>[├│└─┬`\\- ]*)(?P<unmet>UNMET DEPENDENCY )?(?P<name>[^├│└─┬`\\- ].*)@(?P<version>\\S*)(?P<deduped> deduped)?( (?P<addition>.*))?")
	items := treeLineRegex.FindAllStringSubmatch(input, -1)
	parents := make([]*PackageMetadata, 0, 10)
	for _, foundItem := range items {
		treePrefix := foundItem[treeLineRegex.SubexpIndex("treePrefix")]
		indentationLevel := (utf8.RuneCountInString(treePrefix)+1)/2 - 1
		if indentationLevel < 0 {
			indentationLevel = 0
		}
		name := foundItem[treeLineRegex.SubexpIndex("name")]
		version := foundItem[treeLineRegex.SubexpIndex("version")]
		unmet := foundItem[treeLineRegex.SubexpIndex("unmet")] != ""
		newItem := item.New(name)
		newItem.Wanted = version
		if !unmet {
			newItem.Current = version
		}

		if indentationLevel > 0 && len(parents) >= indentationLevel-1 {
			parents = append(parents[0:indentationLevel], newItem)
			parents[indentationLevel-1].Dependencies = append(parents[indentationLevel-1].Dependencies, name)
		} else {
			parents = []*PackageMetadata{newItem}
		}

		result = append(result, newItem)
	}
	var arrayResult []PackageMetadata
	for _, resultItem := range result {
		arrayResult = append(arrayResult, *resultItem)
	}
	return arrayResult, nil
}
