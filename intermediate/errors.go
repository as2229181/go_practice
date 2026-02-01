package main

import (
	"errors"
	"fmt"
)

func validate(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("number is lower than zero")
	}
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("there are no data")
	}
	return nil
}

func main() {
	// result, err := validate(10)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)
	// result, err := validate(-10)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)
	// data_1 := []byte{}
	// if err := process(data_1); err != nil {
	// 	fmt.Println(err)
	// }
	// err_1 := eprocess()
	// if err_1 != nil {
	// 	fmt.Println(err_1)
	// 	return
	// }
	err_2 := readData()
	if err_2 != nil {
		fmt.Println(err_2)
		return
	}

}

type myError struct {
	message string
}

func (m *myError) Error() string{
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Some error happened"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("ReadData: % w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("read config error")
}