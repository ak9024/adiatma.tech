package cmd

import (
	"log"
	"os"

	"github.com/ak9024/adiatma.tech/api/internal/app/server"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "api",
	Long: "CLI to run web server",
	Run: func(cmd *cobra.Command, args []string) {
		if errLoadEnv := godotenv.Load(".env"); errLoadEnv != nil {
			log.Fatal(errLoadEnv)
			os.Exit(1)
		}

		server.Router()
		if err := server.StartApp(os.Getenv("PORT")); err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(cmdEnv)
	if errExecCmd := rootCmd.Execute(); errExecCmd != nil {
		log.Fatal(errExecCmd.Error())
		os.Exit(1)
	}
}
