package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("puzzle-input.txt")
	if err != nil {
		log.Fatalf("could not read puzzle-input.txt file, err=%v", err)
	}
	defer file.Close()

	var freq int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		change, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("could not convert line to integer, err=%v", err)
		}
		freq += change
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("could not read line on frequencies, err=%v", err)
	}

	println(freq)

}
