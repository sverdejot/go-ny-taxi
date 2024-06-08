package cli

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/sverdejot/go-ny-taxi/internal/infrastructure/storage"
)

const (
	filepathFlag = "file"
	hostFlag     = "host"
	portFlag     = "port"
	userFlag     = "user"
	passwordFlag = "pass"
)

func BuildImportCommand() *cobra.Command {
	impt := &cobra.Command{
		Use:   "import",
		Short: "Import NY Taxis records from CSV file",

		Run: func(cmd *cobra.Command, args []string) {
			filePath, _ := cmd.Flags().GetString(filepathFlag)
			host, _ := cmd.Flags().GetString(hostFlag)
			port, _ := cmd.Flags().GetInt(portFlag)
			user, _ := cmd.Flags().GetString(userFlag)
			pass, _ := cmd.Flags().GetString(passwordFlag)

			// just to test out functional options
			connectionString := storage.NewConnectionString(
				storage.WithDriver("postgres"),
				storage.WithHost(host),
				storage.WithPort(port),
				storage.WithUser(user),
				storage.WithPassword(pass),
				storage.WithDatabase("nytaxi"),
				storage.WithOpts(map[string]string{
					"sslmode": "disable",
				}),
			)

			if filePath == "" {
				log.Fatal("no file provided")
			}

			imprt(connectionString.String(), filePath)
		},
	}

	impt.Flags().StringP(filepathFlag, "f", "", "ny taxis csv filepath")
	impt.Flags().StringP(hostFlag, "H", "localhost", "database host")
	impt.Flags().IntP(portFlag, "p", 5432, "database port")
	impt.Flags().StringP(userFlag, "u", "user", "database user")
	impt.Flags().StringP(passwordFlag, "P", "pass", "database pass")

	return impt
}
