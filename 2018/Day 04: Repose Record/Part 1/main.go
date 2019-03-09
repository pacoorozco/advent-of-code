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

	// calculate the amount of time which sleeps each guard
	sleepTimes := map[int]time.Duration{}
	for i, e := range events {
		if e.kind == eventAwake {
			if events[i-1].kind != eventAsleep {
				log.Fatalf("guard #%d awoke from no sleep", e.id)
			}
			sleepTimes[e.id] += e.time.Sub(events[i-1].time)
		}
	}

	// find the guard who sleeps more time
	sleeper := 0
	var maxSleep time.Duration
	for id, d := range sleepTimes {
		if d > maxSleep {
			maxSleep = d
			sleeper = id
		}
	}

	// calculate occurrences for each minute
	minutes := make([]int, 60)
	for i, e := range events {
		if e.id != sleeper || e.kind != eventAwake {
			continue
		}
		for i := events[i-1].time.Minute(); i < e.time.Minute(); i++ {
			minutes[i]++
		}

	}

	// find which minutes has more occurrence
	maxMinute, maxV := 0, minutes[0]
	for i, v := range minutes {
		if v > maxV {
			maxMinute = i
			maxV = v
		}
	}

	return sleeper, maxMinute
}
