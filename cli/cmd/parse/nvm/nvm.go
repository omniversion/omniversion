package nvm

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	. "github.com/omniversion/omniversion/cli/types"
	"github.com/spf13/cobra"
)

func parseNvmOutput(input string) ([]PackageMetadata, error) {
	if isVersionOutput(input) {
		return parseAsVersionOutput(input)
	}

	if isNvmrcFile(input) {
		return parseNvmrcFile(input)
	}

	return parseAsListOutput(input)
}

var ParseCommand = &cobra.Command{
	Use:   "nvm",
	Short: "Parse the output of nvm",
	Long:  `Translate the output of nvm into the omniversion format.`,
	Run:   shared.WrapCommand(parseNvmOutput),
}
