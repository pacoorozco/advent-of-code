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

	var lines []string

	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("could not read line on frequencies, err=%v", err)
	}

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			result := compare(lines[i], lines[j])
			if len(result) == (len(lines[i]) - 1) {
				fmt.Println(result)
			}
		}
	}

}

// compare returns the substring of 'a' that matches with 'b'
func compare(a, b string) string {
	count := ""
	maxIndex := len(a)

	if maxIndex != len(b) {
		return ""
	}

	for i := 0; i < maxIndex; i++ {
			if a[i] == b[i] {
				count = count + string(a[i])
			}
	}

	return count
}
