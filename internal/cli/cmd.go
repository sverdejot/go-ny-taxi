package cli

import (
	"log"

	"github.com/spf13/cobra"
)

const (
	filepathFlag = "file"
)

func BuildImportCommand() *cobra.Command {
	impt := &cobra.Command{
		Use: "import",
		Short: "Import NY Taxis records from CSV file",

		Run: func(cmd *cobra.Command, args []string) {
			filePath, _ := cmd.Flags().GetString(filepathFlag)
			if filePath == "" {
				log.Fatal("no file provided")
			}
			parseFile(filePath)
		},
	}

	impt.Flags().StringP(filepathFlag, "f", "", "ny taxis csv filepath")

	return impt
}
