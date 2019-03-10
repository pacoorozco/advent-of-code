package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := strings.TrimSuffix(string(b), "\n")

	minLength := len(s)
	testedUnits := make(map[string]bool)
	for _, c := range s {
		if testedUnits[strings.ToLower(string(c))] == true {
			continue
		}
		p := removeUnitType(string(c), s)
		length := len(react(p))
		if length < minLength {
			minLength = length
		}
		testedUnits[string(c)] = true
	}

	fmt.Println(minLength)
}

func removeUnitType(unit, polymer string) string {
	polymer = strings.Replace(polymer, strings.ToUpper(unit), "", -1)
	return strings.Replace(polymer, strings.ToLower(unit), "", -1)
}

func react(s string) string {
	ok := true
	for ok {
		s, ok = step(s)
	}
	return s
}

func step(s string) (string, bool) {
	for i := 0; i < len(s)-1; i++ {
		if opposite(s[i], s[i+1]) {
			return s[:i] + s[i+2:], true
		}
	}
	return s, false
}

func opposite(a, b byte) bool {
	const diff = 'a' - 'A'
	return a+diff == b || b+diff == a
}
