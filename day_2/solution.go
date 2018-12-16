package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getInput() []string {

	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	datastr := string(input)
	inputSlice := strings.Split(datastr, "\n")

	return inputSlice
}

func main() {

	input := getInput()

	countExactlyTwoAnyLetter := 0
	countExactlyThreeAnyLetter := 0

	for _, id := range input {
		if hasNumberOfSameCharacters(id, 2) {
			countExactlyTwoAnyLetter++
		}
		if hasNumberOfSameCharacters(id, 3) {
			countExactlyThreeAnyLetter++
		}
	}

	checksum := countExactlyTwoAnyLetter * countExactlyThreeAnyLetter
	fmt.Printf("result: %d\n", checksum)

}

func hasNumberOfSameCharacters(input string, number int) bool {
	for _, char := range input {
		numberOfChar := strings.Count(input, string(char))
		if numberOfChar == number {
			return true
		}
	}
	return false
}
