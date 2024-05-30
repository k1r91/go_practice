package main

import (
	"fmt"
	"errors"
)

type customErr struct {
	desc string
	err error
}

func (ce customErr) Error() string {
	return fmt.Sprintf("%s language error: %v", ce.desc, ce.err)
}

func (ce customErr) Unwrap() error {
	return ce.err
}

func testCustomErr() (string, error) {
	return "", customErr{"test Error", errors.New("some error occured")}
}

func main() {
	_, err := testCustomErr()
	if err != nil {
		fmt.Println(err)
	}
	var ce customErr
	if errors.As(err, &ce) {
		fmt.Println("error is custom error")
	}
}