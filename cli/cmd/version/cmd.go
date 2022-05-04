package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

// Version is the current semantic version
var Version = "0.0.1"

// CommitHash is the long form of the current build's commit hash
var CommitHash = "-"

// RepoChecksum is a hash of all relevant files in the repo at time of build
//
// This will help tell apart binaries built from source by detecting uncommitted changes.
var RepoChecksum = "-"

// Date is the date and time at which the build was created
var Date = "-"

// Via is the installation method ("npm"/"brew"/"compiled from source"/...)
var Via = "compiled from source"

// implements the version command that outputs information on the current installation as a yaml string map
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of the omniversion binary",
	Long:  `All software has versions, even omniversion. This is the current version of your omniversion installation`,
	Run: func(cmd *cobra.Command, args []string) {
		runVersionCmd(cmd.OutOrStdout())
	},
}

func runVersionCmd(writer io.Writer) {
	executableLocation, _ := os.Executable()
	_, _ = writer.Write([]byte(fmt.Sprintln(fmt.Sprintf(
		`name: omniversion
version: %v
via: %v
date: %v
commit: %-8v
checksum: %-8v
location: %v
`, Version, Via, Date, CommitHash, RepoChecksum, executableLocation))))
}
