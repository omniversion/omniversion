package _default

import (
	"fmt"
	"github.com/omniversion/omniversion/cli/cmd/parse/go/helpers"
	"github.com/omniversion/omniversion/cli/cmd/parse/go/item"
	. "github.com/omniversion/omniversion/cli/types"
	"golang.org/x/mod/modfile"
)

func ParseGoModFile(input string) ([]PackageMetadata, error) {
	file, err := modfile.ParseLax("", []byte(input), nil)
	if err != nil {
		return []PackageMetadata{}, fmt.Errorf("unable to parse go.mod file: %q", err)
	}
	var result []PackageMetadata
	if file.Module != nil {
		newItem := item.New(file.Module.Mod.Path)
		newItem.Aliases = []string{helpers.ShortModuleName(file.Module.Mod.Path)}
		newItem.Current = file.Module.Mod.Version
		newItem.PackageManager = "gomod"
		result = append(result, *newItem)
	}
	if file.Go != nil {
		result = append(result, PackageMetadata{
			Name:           "go",
			PackageManager: "go",
			Wanted:         file.Go.Version,
		})
	}
	for _, requireLine := range file.Require {
		if requireLine != nil {
			newItem := item.New(requireLine.Mod.Path)
			newItem.Aliases = []string{helpers.ShortModuleName(requireLine.Mod.Path)}
			newItem.Current = helpers.CleanVersion(requireLine.Mod.Version)
			newItem.PackageManager = "go"
			newItem.InstallPath = requireLine.Mod.Path
			newItem.Homepage = requireLine.Mod.Path
			direct := !requireLine.Indirect
			newItem.Direct = &direct
			result = append(result, *newItem)
		}
	}
	return result, nil
}
