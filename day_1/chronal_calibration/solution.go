package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	input := getInput()

	result, _:= applyChanges(input, 0)

	fmt.Printf("result: %d\n", result)
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

func applyChanges(input []string, start int) (int, bool){


	frequency, result := start, start 
	resultHistory := []int{}

	for _, v := range input {
		change, err := strconv.Atoi(v)

		if err != nil {
			fmt.Println(err)
		}

		result = frequency + change

		fmt.Printf("Current frequency %d, change of %d; resulting frequency %d.\n", frequency, change, result)


		if checkResult(result, resultHistory){

			fmt.Printf("Duplicate result found: %d", result)
			return result, true

		} else{
			resultHistory = append(resultHistory, result)
		}

		frequency = result
	}

	return result, false
}


func checkResult(result int, history[]int)bool{

	for _, item := range history{
		if item == result{
			return true
		}
	}
	return false
}