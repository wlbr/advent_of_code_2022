package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(fname string) (monkeys []*monkey) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	tmpmonkeys := make(map[int]*monkey)

	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		m := monkey{}

		line := scanner.Text()
		if line == "" {
			scanner.Scan()
			line = scanner.Text()
		}
		_, end := fmt.Sscanf(line, "Monkey %d:\n", &m.id)
		if end != nil {
			log.Fatal("Error parsing monkey)")
		}

		scanner.Scan()
		line = scanner.Text()
		for _, item := range strings.Split(strings.Split(line, ":")[1], ",") {
			iitem, err := strconv.Atoi(strings.Trim(item, " "))
			if err != nil {
				log.Fatalf("Error parsing item '%s': %s", item, err)
			}
			m.items = append(m.items, iitem)
		}

		scanner.Scan()
		line = scanner.Text()
		var ops, arg string
		_, end = fmt.Sscanf(line, "  Operation: new = old %s %s\n", &ops, &arg)
		if end != nil {
			log.Fatalf("Error parsing monkey: %s", end)
		}
		m.operation = NewOperation(ops, arg)

		scanner.Scan()
		line = scanner.Text()
		var dec string
		var decarg, truetarget, falsetarget int
		_, end = fmt.Sscanf(line, "  Test: %s by %d\n", &dec, &decarg)
		if end != nil {
			log.Fatal("Error parsing monkey", end)
		}

		scanner.Scan()
		line = scanner.Text()
		_, end = fmt.Sscanf(line, "  If true: throw to monkey %d\n", &truetarget)
		if end != nil {
			log.Fatal("Error parsing monkey", end)
		}

		scanner.Scan()
		line = scanner.Text()
		_, end = fmt.Sscanf(line, "  If false: throw to monkey %d\n", &falsetarget)
		if end != nil {
			log.Fatal("Error parsing monkey)")
		}
		m.decision = NewDecision(dec, decarg, truetarget, falsetarget)
		tmpmonkeys[m.id] = &m
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	max := 0
	for _, m := range tmpmonkeys {
		if m.id > max {
			max = m.id
		}
	}
	monkeys = make([]*monkey, max+1)
	for _, m := range tmpmonkeys {
		monkeys[m.id] = m
	}

	return monkeys
}
