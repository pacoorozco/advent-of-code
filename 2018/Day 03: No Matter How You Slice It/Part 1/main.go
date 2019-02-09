package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("could not read input data from file, err=%v", err)
	}
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	var f fabric

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var id, x, y, w, h int
		_, err = fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatal(err)
		}
		f.addClaim(id, x, y, w, h)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("could not read line on frequencies, err=%v", err)
	}

	fmt.Println(f.claimedMoreThanOne())

}

type xy struct { x, y int }

type fabric struct {
	coord map[xy] int
}

func (f *fabric) addClaim(id, x, y, w, h int) {
	if f.coord == nil {
		f.coord = make(map[xy]int)
	}

	for i := x; i < (x + w); i++ {
		for j := y; j < (y + h); j++ {
			f.coord[xy{i,j}]++
		}
	}
}

func (f *fabric) claimedMoreThanOne() int {
	count := 0;
	for _, p := range f.coord {
		if p > 1 {
			count++
		}
	}
	return count
}

func (f *fabric) Print() {
	var maxX, maxY int

	for p := range f.coord {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			fmt.Print(f.coord[xy{x,y}])
		}
		fmt.Println()
	}
}
