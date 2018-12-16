package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	input := getInput()

	total := 0
	for _, v := range input {
		i, err := strconv.Atoi(v)

		if err != nil {
			fmt.Println(err)
		}
		total += i
	}

	fmt.Printf("result: %d\n", total)

}

func getInput() []string {

	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	datastr := string(input)
	inputSlice := strings.Split(datastr, "\n")

	return inputSlice
}
