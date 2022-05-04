package cmd

import (
	. "github.com/omniversion/omniversion/cli/cmd/parse"
	"github.com/omniversion/omniversion/cli/cmd/parse/apt"
	"github.com/omniversion/omniversion/cli/cmd/parse/npm"
	"github.com/omniversion/omniversion/cli/cmd/parse/rubygems"
	"github.com/omniversion/omniversion/cli/cmd/parse/rvm"
	. "github.com/omniversion/omniversion/cli/cmd/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var log = logrus.New()

var RootCmd = &cobra.Command{
	Use:   "omniversion",
	Short: "omniversion",
	Long:  `Manage versions, dependencies and vulnerabilities`,
}

func init() {
	RootCmd.AddCommand(ParseCmd)
	ParseCmd.AddCommand(apt.ParseCommand)
	ParseCmd.AddCommand(npm.ParseCommand)
	ParseCmd.AddCommand(rubygems.ParseCommand)
	ParseCmd.AddCommand(rvm.ParseCommand)
	RootCmd.AddCommand(VersionCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
