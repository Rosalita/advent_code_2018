package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"math"
)

type coord struct{
	x int
	y int
}

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
	//input := getInput() 

	// create a grid
	// for each location on the grid, find the closest coord
	// if more more than 1 coord is closest don't count as closest to any

	// to find the distance between two coords
	// p1, p2 and q1, q2
    // abs(p1 - q1) + abs(p2 - q2)  

	coord1 := coord{
		x: 0,
		y: 0,
	}

	coord2 := coord{
		x: 1,
		y: 6,
	}

	fmt.Printf("distance is %d\n", distance(coord1, coord2))
}

func distance(a, b coord) int{

	xAbs := math.Abs(float64(a.x) - float64(b.x))
	yAbs := math.Abs(float64(a.y) - float64(b.y))

	return int(xAbs + yAbs)
}