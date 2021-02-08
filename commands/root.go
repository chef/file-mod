package commands

import (
	"log"

	"github.com/spf13/cobra"

	filesystem "github.com/chef/go-filesystem"
)

var (
	version string

	fs filesystem.FileSystem

	rootCmd = &cobra.Command{
		Use:          "file-mod",
		Short:        "Command line utility to modify files.",
		SilenceUsage: true,
		Version:      version,
	}
)

func init() {
	rootCmd.SetVersionTemplate("{{.Version}}\n")
}

// Execute handles the execution of child commands and flags.
func Execute() {
	fs = filesystem.NewOsFs()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
