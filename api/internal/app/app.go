package app

func StartApp() error {
	r := Router()

	if err := r.Listen(":8000"); err != nil {
		return err
	}

	return nil
}
