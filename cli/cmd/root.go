package cmd

import (
	"github.com/omniversion/omniversion/cli/cmd/parse"
	"github.com/omniversion/omniversion/cli/cmd/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var log = logrus.New()

var Cmd = &cobra.Command{
	Use:   "omniversion",
	Short: "omniversion",
	Long:  `Manage versions, dependencies and vulnerabilities`,
}

func init() {
	parse.InitSubcommand(Cmd)
	version.InitSubcommand(Cmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
