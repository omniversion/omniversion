package homebrew

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
)

func parseHomebrewOutput(input string) ([]PackageMetadata, error) {
	if isVersionCommandOutput(input) {
		return parseAsVersionCommandOutput(input)
	}

	if isDefaultListOutput(input) {
		return parseAsDefaultListOutput(input)
	}

	return parseAsListWithVersionsCommandOutput(input)
}

var ParseCommand = &cobra.Command{
	Use:     "homebrew",
	Short:   "Parse the output of homebrew",
	Long:    `Translate the output of homebrew into the omniversion format.`,
	Run:     shared.WrapCommand(parseHomebrewOutput),
	Aliases: []string{"brew"},
}
