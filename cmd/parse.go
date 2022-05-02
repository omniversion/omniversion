package cmd

import (
	"github.com/omniversion/omniversion-cli/cmd/parse"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse the output of the specified package manager",
	Long:  `Transform the output of a package manager into a common format. You will need to specify the package manager in question.`,
}

func init() {
	rootCmd.AddCommand(parseCmd)
	parseCmd.AddCommand(parse.AptCmd)
	parseCmd.AddCommand(parse.NpmCmd)
	parseCmd.AddCommand(parse.RubygemsCmd)
	parseCmd.AddCommand(parse.RvmCmd)
}
