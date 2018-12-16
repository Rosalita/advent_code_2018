package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type claim struct {
	x      int
	y      int
	width  int
	height int
}

func getInput() []claim {

	claims := []claim{}

	input, err := ioutil.ReadFile("small_input.txt")

	if err != nil {
		fmt.Println(err)
	}

	datastr := string(input)
	inputSlice := strings.Split(datastr, "\n")

	for i := range inputSlice {
		splitString := strings.Split(inputSlice[i], "@")
		noID := strings.TrimSpace(splitString[1])
		coordAndSize := strings.Split(noID, ":")
		coord := coordAndSize[0]
		size := strings.TrimSpace(coordAndSize[1])
		y, x := splitToInt(coord, ",")
		w, h := splitToInt(size, "x")

		claim := claim{
			x:      x,
			y:      y,
			width:  w,
			height: h,
		}

		claims = append(claims, claim)

	}

	return claims
}

func main() {
	input := getInput()

	fabric := make([][]int, 8)
	for i := range fabric {
		fabric[i] = make([]int, 8)
	}

	for _, claim := range input {
		addClaim(fabric, claim)
	}

	for _, row := range fabric{
		fmt.Println(row)
	}

	result := countInches(fabric)
	fmt.Printf("result is: %d\n",result)

}

func splitToInt(str string, sep string) (a int, b int) {
	split := strings.Split(str, sep)
	a, _ = strconv.Atoi(split[0])
	b, _ = strconv.Atoi(split[1])
	return a, b
}

func addClaim(fabric [][]int, c claim) [][]int {
	for i := 0; i < c.width; i++ {
		for j := 0; j < c.height; j++ {
			fabric[i+c.x][j+c.y]++
		}
	}
	return fabric
}

func countInches(fabric [][]int) int {

	count := 0

	for _, row := range fabric {
		for _, value := range row {
			if value > 1 {
				count++
			}
		}
	}

	return count
}
