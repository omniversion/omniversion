package parse

import (
	"github.com/spf13/cobra"
)

var ParseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse the output of the specified package manager",
	Long:  `Transform the output of a package manager into a common format. You will need to specify the package manager in question.`,
}
