package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"strconv"
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

	shifts := splitIntoShifts(input)
	fmt.Println(shifts)

	//guardsSleepingMins := map[string]int{}
	

	for _, shift := range shifts{
				//"#971", "asleep15", "wakes54",

		guard := ""
		minutes := 0
		sleepAt := 0
		wokeAt := 0

		for i, event := range shift{

			if i == 0{
				guard = event
				continue
			}

			if i % 2 == 0 { // if even index event
				mins:= event[len(event)-2:]
				sleepAt, _ = strconv.Atoi(mins)
			}
			if i % 2 == 1{ // if odd index event
				mins := event[len(event)-2:]
				wokeAt, _ = strconv.Atoi(mins)

				timeAsleep := wokeAt - sleepAt
				minutes += timeAsleep
			}
		}
	
		fmt.Println(guard)
		fmt.Println(minutes)
		//"#971", "asleep15", "wakes54",
		//"#3079", "asleep08", "wakes48", "asleep51", "wakes52",


	}

}

func splitIntoShifts(input []string) [][]string {
	shifts := [][]string{}
	shift := []string{}

	for i, row := range input {

		minutes := row[15:17]

		if i == 0 {
			splitstring := strings.Split(row, " ")
			shift = append(shift, splitstring[3])
			continue
		}

		if i == (len(input) - 1) {
			if strings.Contains(row, "asleep") {
				shift = append(shift, "asleep"+minutes)
			} else {
				shift = append(shift, "wakes"+minutes)
			}
			shifts = append(shifts, shift)
		}

		if strings.Contains(row, "Guard") {
			shifts = append(shifts, shift)
			shift = []string{}
			splitstring := strings.Split(row, " ")
			shift = append(shift, splitstring[3])

		} else {
			if strings.Contains(row, "asleep") {
				shift = append(shift, "asleep"+minutes)
			} else {
				shift = append(shift, "wakes"+minutes)
			}
		}
	}

	return shifts
}
