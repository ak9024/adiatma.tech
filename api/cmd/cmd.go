package cmd

import (
	"log"
	"os"

	"github.com/ak9024/adiatma.tech/api/internal/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {
		app.Router()

		if err := app.StartApp(); err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
