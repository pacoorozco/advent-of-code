package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type guardEvent struct {
	id   int
	kind eventKind
	time time.Time
}

func (e guardEvent) String() string {
	date := e.time.Format("01/02 15:04")
	switch e.kind {
	case eventStart:
		return fmt.Sprintf("[%s] Guard #%d starts", date, e.id)
	case eventAsleep:
		return fmt.Sprintf("[%s] Guard #%d sleeps", date, e.id)
	case eventAwake:
		return fmt.Sprintf("[%s] Guard #%d wakes", date, e.id)
	}
	return fmt.Sprintf("unknown guard event type: %#v", e)
}

type eventKind byte

const (
	eventStart eventKind = iota
	eventAsleep
	eventAwake
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	sort.Strings(lines)

	var events []guardEvent
	var currentGuard int

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		dateEnd := strings.Index(line, "]")
		dateText := line[1:dateEnd]
		date, err := time.Parse("2006-01-02 15:04", dateText)
		if err != nil {
			log.Fatalf("could not parse date %s: %v", dateText, err)
		}
		e := guardEvent{time: date}
		pieces := strings.Fields(line[dateEnd+2:])
		switch pieces[0] {
		case "Guard":
			id, err := strconv.Atoi(pieces[1][1:])
			if err != nil {
				log.Fatalf("could not parse id %s: %v", pieces[1][1:], err)
			}
			currentGuard = id
			e.id = id
			e.kind = eventStart
		case "falls":
			e.id = currentGuard
			e.kind = eventAsleep
		case "wakes":
			e.id = currentGuard
			e.kind = eventAwake
		}

		events = append(events, e)
	}

	id, minute := findGuard(events)
	fmt.Println(id * minute)
}

func findGuard(events []guardEvent) (id, minute int) {

	minutes := make([]map[int]int, 60)
	for i := range minutes {
		minutes[i] = make(map[int]int)
	}

	for i, e := range events {
		if e.kind == eventAwake {
			if events[i-1].kind != eventAsleep {
				log.Fatalf("guard #%d awoke from no sleep", e.id)
			}
			for m := events[i-1].time.Minute(); m < e.time.Minute(); m++ {
				minutes[m][e.id]++
			}
		}
	}

	maxMinute, maxID, maxCount := 0, 0, 0

	for minute, counts := range minutes {
		for id, n := range counts {
			if n > maxCount {
				maxMinute = minute
				maxCount = n
				maxID = id
			}
		}
	}

	return maxID, maxMinute
}
