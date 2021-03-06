package parse

import (
	"github.com/omniversion/omniversion/cli/cmd/parse/apt"
	"github.com/omniversion/omniversion/cli/cmd/parse/brew"
	"github.com/omniversion/omniversion/cli/cmd/parse/galaxy"
	"github.com/omniversion/omniversion/cli/cmd/parse/gem"
	_go "github.com/omniversion/omniversion/cli/cmd/parse/go"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm"
	"github.com/omniversion/omniversion/cli/cmd/parse/nvm"
	"github.com/omniversion/omniversion/cli/cmd/parse/pip"
	"github.com/omniversion/omniversion/cli/cmd/parse/raw"
	"github.com/omniversion/omniversion/cli/cmd/parse/rvm"
	"github.com/omniversion/omniversion/cli/cmd/parse/shared"
	"github.com/spf13/cobra"
)

func InitSubcommand(rootCommand *cobra.Command) {
	Cmd.PersistentFlags().BoolVarP(&shared.InjectPackageManager, "inject-pm", "i", false, "inject package manager name into each package metadata item")
	Cmd.PersistentFlags().StringVarP(&shared.OutputFormat, "output-format", "o", "yaml", "the desired output format (\"toml\"/\"json\"/\"yaml\")")

	rootCommand.AddCommand(Cmd)
	Cmd.AddCommand(galaxy.ParseCommand)
	Cmd.AddCommand(apt.ParseCommand)
	Cmd.AddCommand(_go.ParseCommand)
	Cmd.AddCommand(brew.ParseCommand)
	Cmd.AddCommand(npm.ParseCommand)
	Cmd.AddCommand(nvm.ParseCommand)
	Cmd.AddCommand(pip.ParseCommand)
	Cmd.AddCommand(raw.ParseCommand)
	Cmd.AddCommand(gem.ParseCommand)
	Cmd.AddCommand(rvm.ParseCommand)
}

var Cmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse the output of the specified package manager",
	Long:  `Transform the output of a package manager into a common format. You will need to specify the package manager in question.`,
}
