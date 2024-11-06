package main

import "fmt"

func main() {
	fmt.Printf("%T type error", f("w"))
}

func producer(s string) error {
	if s == "w" {
		return fmt.Errorf("w error") // *fmt.wrapError
	}
	if s == "v" {
		return fmt.Errorf("v error") // *errors.errorString
	}
	return nil
}

func f(s string) error {
	if err := producer(s); err != nil {
		return fmt.Errorf("producer call error %v", err)
	}
	return nil
}

// legacy way of wrapping an error
// was to create a special type
// for that specific error.

// more flexible error management
type BarError struct {
	Err error
}

// implement error interface
func (b *BarError) Error() string {
	return "bar failed: " + b.Err.Error()
}
