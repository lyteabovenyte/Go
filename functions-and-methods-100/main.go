package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type MultiError struct {
	err []string
}

type Customer struct {
	Name string
	Age  int
}

// add an error to the struct
func (m *MultiError) Add(err error) {
	m.err = append(m.err, err.Error())
}

// implementing error interface
func (m *MultiError) Error() string {
	return strings.Join(m.err, ";")
}

// returning an interface
func (c Customer) validate() error {
	var m *MultiError // nil pointer value.
	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("customer age cannot be less than zero"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("customer name cannot be empty"))
	}
	if m == nil {
		return nil // returning nil value explicitly
	}
	return m
}

func main() {

	fmt.Println("-----------Reader experiment---------")
	reader := strings.NewReader("more clever, more smart. use your brain")
	p := make([]byte, 39)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
	fmt.Println("-----------Reader experiment---------")

	c := Customer{Name: "", Age: 24}
	if err := c.validate(); err != nil {
		fmt.Printf("customer invalid: %v\n", err)
	} else {
		fmt.Println("customer is valid.")
	}
}
