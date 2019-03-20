package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := GetLines("../input.txt")

	x := make([] int, len(lines))
	y := make([] int, len(lines))

	bigX, bigY := 0, 0

	for n, v := range lines {
		split := RegSplit(v, "[, ]+")
		i, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatalf("problem converting coordinate x=%s, %v", split[0], err)
		}
		x[n] = i

		if i > bigX {
			bigX = i
		}

		j, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("problem converting coordinate y=%s, %v", split[1], err)
		}
		y[n] = j

		if j > bigY {
			bigY = j
		}
	}

	grid := make([][]int, bigX+1)
	for i := range grid {
		grid[i] = make([] int, bigY+1)
	}

	counted := make([]int, len(x))
	for i := 0; i < bigX+1; i++ {
		for j := 0; j < bigY+1; j++ {
			c := closest(i, j, x, y)
			grid[i][j] = c
			if c != -1 && counted[c] != -1 {
				if i == bigX || i == 0 || j == 0 || j == bigY {
					counted[c] = -1
				} else {
					counted[c]++
				}
			}
		}
	}
	fmt.Println(counted)

	// find the maximum in counted array
	value := 0
	for i := 0; i < len(counted); i++ {
		if counted[i] > value {
			value = counted[i]
		}
	}

	fmt.Println(value)

}

func closest(x, y int, x1, y1 []int) int {
	closest, distance := 0, math.MaxInt32
	equal := false
	for n := range y1 {
		man := manhattanDistance(x, y, x1[n], y1[n])

		if man < distance {
			distance = man
			closest = n
			equal = false
		} else if man == distance {
			equal = true
		}
	}
	if equal {
		return -1
	}
	return closest
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// GetLines gets the lines from a file
func GetLines(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// RegSplit A simple regex splitter
func RegSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:]
	return result
}
