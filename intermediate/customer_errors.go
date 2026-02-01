package main

import (
	"errors"
	"fmt"
)

func main() {
	err := dosomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("nothing happened")
}


type CustomError struct {
	code    int
	message string
	err error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, error: %v \n", e.code, e.message, e.err)
}

func dosomething() error {
	if err := doSomethingElse(); err != nil{
		return &CustomError{
			code: 400,
			message: "something went wrong",
			err: err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("Internal error")
}