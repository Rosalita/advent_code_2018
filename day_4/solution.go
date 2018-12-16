package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
)


func getInput() []string {


	input, err := ioutil.ReadFile("small_input.txt")

	if err != nil {
		fmt.Println(err)
	}

	datastr := string(input)
	inputRows := strings.Split(datastr, "\n")

	return inputRows
}

func main() {
	input := getInput()
	sort.Strings(input)

	shifts := [][]string{}
	shift := []string{}

	for _, row := range input{

		splitstring := strings.Split(row, " ")
		minutes := row[15:17]
		if strings.Contains(row, "Guard"){
			shifts = append(shifts, shift)
			shift = []string{}
			shift = append(shift, splitstring[3])
		}

		if strings.Contains(row, "asleep"){
			shift = append(shift, "asleep" + minutes)
		}
		if strings.Contains(row, "wakes"){
			shift = append(shift, "wakes" + minutes)
		}

	}

	fmt.Printf("%+v\n", shifts)

	

}
