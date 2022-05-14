package _go

import (
	"fmt"
	"github.com/omniversion/omniversion/cli/cmd/parse/go/helpers"
	"github.com/omniversion/omniversion/cli/cmd/parse/go/item"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
)

func ParseGoModFile(input string) ([]PackageMetadata, error) {
	file, err := modfile.ParseLax("", []byte(input), nil)
	if err != nil {
		return []PackageMetadata{}, fmt.Errorf("unable to parse go.mod file: %q", err)
	}
	var result []PackageMetadata
	if file.Module != nil {
		newItem := item.New(helpers.NameForPath(file.Module.Mod.Path))
		newItem.Current = file.Module.Mod.Version
		newItem.PackageManager = "gomod"
		newItem.InstallPath = file.Module.Mod.Path
		result = append(result, *newItem)
	}
	if file.Go != nil {
		result = append(result, PackageMetadata{
			Name:           "go",
			PackageManager: "gomod",
			Wanted:         file.Go.Version,
		})
	}
	for _, requireLine := range file.Require {
		if requireLine != nil {
			newItem := item.New(helpers.NameForPath(requireLine.Mod.Path))
			newItem.Current = helpers.CleanVersion(requireLine.Mod.Version)
			newItem.PackageManager = "gomod"
			newItem.InstallPath = requireLine.Mod.Path
			newItem.Homepage = requireLine.Mod.Path
			direct := !requireLine.Indirect
			newItem.Direct = &direct
			result = append(result, *newItem)
		}
	}
	return result, nil
}

func parseOutput(input string) ([]PackageMetadata, error) {
	return ParseGoModFile(input)
}

var ParseCommand = &cobra.Command{
	Use:     "gomod",
	Short:   "Parse the output of go mod",
	Long:    `Translate the output of npm into the omniversion format.`,
	Run:     shared.WrapCommand(parseOutput),
	Aliases: []string{"go", "mod"},
}
