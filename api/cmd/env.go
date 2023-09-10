package cmd

import (
	"log"
	"os"

	"github.com/ak9024/adiatma.tech/api/internal/app/env"
	"github.com/spf13/cobra"
)

var cmdEnv = &cobra.Command{
	Use:  "env",
	Long: "CLI to generate env",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := ".env"
		// generate file .env base on os environment
		errGenerateEnv := env.GenerateEnv(fileName)
		if errGenerateEnv != nil {
			log.Fatal(errGenerateEnv.Error())
			os.Exit(1)
		}
	},
}
