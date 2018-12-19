package main

import (
	"fmt"
	"io/ioutil"
)

func getInput() string {

	input, err := ioutil.ReadFile("small_input.txt")

	if err != nil {
		fmt.Println(err)
	}

	return string(input)
}

func main() {

	input := getInput()
	fmt.Println(input)

	for _, char := range input {
		fmt.Printf("%v %v\n",char, string(char))

		// a = 97 (-32 = A)
		// A = 65 (+ 32 = a) 

		// val1, val2 = val1 - val2 == 32 or -32
		// make new slice without those two values
	}
}
