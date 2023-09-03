package env

import (
	"fmt"
	"os"
)

func GenerateEnv(fileName string) error {
	port := os.Getenv("PORT")

	mapEnv := map[string]interface{}{
		"PORT": port,
	}

	file, errCreateFile := os.Create(fileName)
	if errCreateFile != nil {
		return errCreateFile
	}

	for key, value := range mapEnv {
		_, errGenerateEnv := fmt.Fprintf(file, "%s=%s\n", key, value)
		if errGenerateEnv != nil {
			return errGenerateEnv
		}
	}

	defer file.Close()

	return nil
}
