package env

import (
	"fmt"
	"os"
)

func GenerateEnv(fileName string) error {
	// get data from os environment
	port := os.Getenv("PORT")
	basePath := os.Getenv("BASE_PATH")
	host := os.Getenv("HOST")

	// map the env
	mapEnv := map[string]interface{}{
		"PORT":      port,
		"BASE_PATH": basePath,
		"HOST":      host,
	}

	// create file name
	file, errCreateFile := os.Create(fileName)
	if errCreateFile != nil {
		return errCreateFile
	}

	for key, value := range mapEnv {
		// insert data to file
		_, errGenerateEnv := fmt.Fprintf(file, "%s=%s\n", key, value)
		if errGenerateEnv != nil {
			return errGenerateEnv
		}
	}

	defer file.Close()

	return nil
}
