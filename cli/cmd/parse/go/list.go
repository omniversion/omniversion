package _go

import (
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/go/helpers"
	"github.com/omniversion/omniversion/cli/cmd/parse/go/item"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
)

func ParseListOutput(input string) ([]PackageMetadata, error) {
	extractionRegex := regexp.MustCompile(`(?m)^(?P<path>\S+)( (?P<versions>.+))?$`)
	items := extractionRegex.FindAllStringSubmatch(input, -1)

	result := make([]PackageMetadata, 0, len(items))
	var allErrors *multierror.Error
	for _, itemData := range items {
		path := itemData[extractionRegex.SubexpIndex("path")]
		versions := itemData[extractionRegex.SubexpIndex("versions")]
		newItem := item.New(helpers.NameForPath(path))
		currentVersion := helpers.LastVersion(strings.Split(versions, " "))
		newItem.Current = currentVersion
		newItem.InstallPath = path
		newItem.Installations = []InstalledPackage{{
			Version: currentVersion,
		}}
		if strings.Count(versions, " ") > 0 {
			newItem.Sources = []PackagesSource{{
				Versions: helpers.CleanVersions(strings.Split(versions, " ")),
			}}
		}
		result = append(result, *newItem)
	}
	return result, allErrors.ErrorOrNil()
}
