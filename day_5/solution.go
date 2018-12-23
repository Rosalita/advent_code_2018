package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

func getInput() string {

	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	return string(input)
}

func main() {

	input := getInput()

	result, done := reduce(input)

	for !done {
		result, done = reduce(result)
	}
	
	fmt.Printf("result is : %s\n", result)
	fmt.Printf("%d units remain\n", len(result))
}


func reduce(input string)(result string, done bool){

	for i := range input {
		if i == 0 {
			continue
		}

		difference := float64(input[i-1]) - float64(input[i])
		absDifference := math.Abs(difference)

		if absDifference == 32{
			firstHalfString := input[:i-1]
			secondHalfString := input[i+1:]

			returnString := firstHalfString + secondHalfString
			return returnString, false
		}
	}

	return input, true
}