package json

import (
	"encoding/json"
	_errors "errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/omniversion/omniversion/cli/cmd/parse/go/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"regexp"
	"strings"
	"time"
)

type ListJsonOutput struct {
	Path      string
	Main      bool
	Dir       string
	GoMod     string
	GoVersion string
	Version   string
	Indirect  bool
	Update    *struct {
		Path    string
		Version string
		Time    time.Time
	}
	Time time.Time
}

type ModuleError struct {
	Err string // the error itself
}

func ParseListOutput(input string) ([]PackageMetadata, error) {
	// `go list -m -json` returns a concatenated string of JSON objects instead of a valid array
	jsonObjectFinder := regexp.MustCompile(`(?m)^{(.|\n)*?\n}`)

	var result []PackageMetadata
	var errors *multierror.Error
	for _, jsonObjectString := range jsonObjectFinder.FindAllString(input, -1) {
		var listItem ListJsonOutput
		jsonUnmarshallErr := json.Unmarshal([]byte(jsonObjectString), &listItem)
		if jsonUnmarshallErr != nil {
			// maybe it's a `ModuleError` instead?
			var moduleError ModuleError
			jsonErrorUnmarshallErr := json.Unmarshal([]byte(jsonObjectString), &moduleError)
			if jsonErrorUnmarshallErr == nil {
				errors = multierror.Append(errors, _errors.New(moduleError.Err))
			} else {
				errors = multierror.Append(errors, fmt.Errorf("failed to parse JSON: %q", jsonUnmarshallErr))
			}
			continue
		}

		newItem := item.New(listItem.Path)
		newItem.Aliases = []string{shared.ShortModuleName(listItem.Path)}
		newItem.Current = shared.CleanVersion(listItem.Version)
		direct := !listItem.Indirect
		newItem.Direct = &direct
		if listItem.Dir != "" {
			version := ""
			if strings.Contains(listItem.Dir, "@") {
				version = shared.LastVersion(strings.Split(listItem.Dir, "@"))
			}
			newItem.Installations = []InstalledPackage{{
				Location: listItem.Dir,
				Version:  version,
			}}
		}
		if listItem.Update != nil {
			newItem.Latest = shared.CleanVersion(listItem.Update.Version)
			newItem.Sources = []PackagesSource{{
				Versions: []string{shared.CleanVersion(listItem.Update.Version)},
			}}
		}
		result = append(result, *newItem)
	}
	return result, errors.ErrorOrNil()
}
