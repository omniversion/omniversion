package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var log = logrus.New()

var RootCmd = &cobra.Command{
	Use:   "omniversion",
	Short: "omniversion",
	Long:  `Manage versions, dependencies and vulnerabilities`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
