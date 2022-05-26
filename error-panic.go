package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// strconv.Atoi("123") => merubah string menjadi number

	var input string
	fmt.Println("Type smoe number: ")
	fmt.Scanln(&input)

	var number int
	var err error

	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number)
	} else {
		panic(err.Error())
	}

	// if valid, err := validate(input); valid {
	// 	fmt.Println(input)
	// } else {
	// 	fmt.Println("Error:", err.Error())
	// }
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" { // trim space => menghilangkan spasi sesudah dan sebelum string
		return false, errors.New("Input is empty") // merubah pesan error
	}
	return true, nil
}
