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

	scanner := bufio.NewScanner(file)

	countThree := 0
	countFour := 0

	for scanner.Scan() {
		m := make(map[rune] int)
		line := scanner.Text()

		for _, char := range line {
			m[char] = m[char] + 1
		}

		notThree := true
		notFour := true

		for _, count := range m {
			if count == 2 && notThree {
				notThree = false
				countThree++
			}
			if count == 3 && notFour {
				notFour = false
				countFour++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("could not read line on frequencies, err=%v", err)
	}

	checksum := countFour * countThree
	fmt.Println(checksum)
}
