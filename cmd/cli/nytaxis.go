package main

import (
	"github.com/spf13/cobra"
	"github.com/sverdejot/go-ny-taxi/internal/cli"
)

func main() {
	rootCmd := &cobra.Command{Use: "import"}
	rootCmd.AddCommand(cli.BuildImportCommand())

	rootCmd.Execute()
}
