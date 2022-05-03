package parse

import (
	root "github.com/omniversion/omniversion/cli/cmd"
	"github.com/omniversion/omniversion/cli/cmd/parse/apt"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm"
	"github.com/omniversion/omniversion/cli/cmd/parse/rubygems"
	"github.com/omniversion/omniversion/cli/cmd/parse/rvm"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse the output of the specified package manager",
	Long:  `Transform the output of a package manager into a common format. You will need to specify the package manager in question.`,
}

func init() {
	root.RootCmd.AddCommand(parseCmd)
	parseCmd.AddCommand(apt.ParseCommand)
	parseCmd.AddCommand(npm.ParseCommand)
	parseCmd.AddCommand(rubygems.ParseCommand)
	parseCmd.AddCommand(rvm.ParseCommand)
}
