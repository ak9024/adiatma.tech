package server

import "fmt"

func StartApp(port string) error {
	r := Router()

	// @TODO need to refactor to implement gracefull shutdown if run in production mode
	// starting the app
	if err := r.Listen(fmt.Sprintf(":%s", port)); err != nil {
		return err
	}

	return nil
}
