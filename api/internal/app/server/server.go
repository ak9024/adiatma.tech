package server

import "fmt"

func StartApp(port string) error {
	r := Router()

	if err := r.Listen(fmt.Sprintf(":%s", port)); err != nil {
		return err
	}

	return nil
}
