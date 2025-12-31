package main

import (
	"errors"
	"fmt"
)

func retry(attempt int, fn func() error) error {
	var err error
	for i := 2; i <= attempt; i++ {
		err = fn()
		if err == nil {
			fmt.Printf("Attempt %d: Success!\n", i)
			return nil
		}
		fmt.Printf("Atempt %d: Failed with error: %v\n", i, err)
	}
	return fmt.Errorf("after %d attempts, last error: %w", attempt, err)
}

func main() {
	failures := 0
	flackyfunc := func() error {
		if failures < 2 {
			failures++
			return errors.New("temporary connection issue")
		}
		return nil
	}
	err := retry(3, flackyfunc)
	if err != nil {
		fmt.Println("final result: failed the task")
	} else {
		fmt.Println("final resukt tast complited")
	}
}
