package Book

func retry(attempt int, fn func() error) error {
	for i := 0; i < attempt; i++ {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
