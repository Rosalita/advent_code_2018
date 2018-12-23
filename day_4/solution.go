package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func getInput() []string {

	input, err := ioutil.ReadFile("input.txt")

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
	guardsSleepingMins := map[string]int{}

	for _, shift := range shifts {

		guard := ""
		minutes := 0
		sleepAt := 0
		wokeAt := 0

		for i, event := range shift {

			if i == 0 {
				guard = event
				continue
			}

			if i%2 == 1 { // if odd index event
				mins := event[len(event)-2:]
				sleepAt, _ = strconv.Atoi(mins)
			}
			if i%2 == 0 { // if even index event
				mins := event[len(event)-2:]
				wokeAt, _ = strconv.Atoi(mins)

				timeAsleep := wokeAt - sleepAt
				minutes += timeAsleep
			}
		}
		guardsSleepingMins[guard] += minutes

	}

	guardNumber, max := max(guardsSleepingMins)

	fmt.Printf("The guard that has slept the most minutes is %s %d\n", guardNumber, max)

	sleepiestGuardShifts := [][]string{}

	for _, shift := range shifts {

		if shift[0] == guardNumber {
			sleepiestGuardShifts = append(sleepiestGuardShifts, shift)
		}

	}

	fmt.Println(sleepiestGuardShifts)
	SleepFreq := map[int]int{}

	for _, shift := range sleepiestGuardShifts {

		sleepAt := 0
		wokeAt := 0



		for i, event := range shift {

			if i == 0 {
				continue
			}

			if i%2 == 1 { // if odd index event
				mins := event[len(event)-2:]
				sleepAt, _ = strconv.Atoi(mins)
				fmt.Printf("asleep at %d\n", sleepAt)
			}
			if i%2 == 0 { // if even index event
				mins := event[len(event)-2:]
				wokeAt, _ = strconv.Atoi(mins)
				fmt.Printf("woke at %d\n", wokeAt)


				for j := sleepAt; j < wokeAt; j++ {
					SleepFreq[j] ++
				}
			}
		}

	}

	fmt.Println(SleepFreq)

	minuteNum, max := maxInt(SleepFreq)

	fmt.Printf("the guard spent minute %d asleep for %d minutes in total\n", minuteNum, max)

	fmt.Println(971 * 38)

}

func max(results map[string]int) (string, int) {
	var max int
	var guardNumber string
	for guard, minutes := range results {
		max = minutes
		guardNumber = guard
		break
	}
	for guard, minutes := range results {
		if minutes > max {
			max = minutes
			guardNumber = guard
		}
	}
	return guardNumber, max
}

func maxInt(results map[int]int) (int, int) {
	var max int
	var minuteNum int
	for minuteNumber, freq := range results {
		max = freq
		minuteNum = minuteNumber
		break
	}
	for minuteNumber, freq := range results {
		if freq > max {
			max = freq
			minuteNum = minuteNumber
		}
	}
	return minuteNum, max
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
