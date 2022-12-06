package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type sortedgroup struct {
	cals      []int
	groupsize int
}

func (m *sortedgroup) add(c int) {
	if len(m.cals) < m.groupsize {
		m.cals = append(m.cals, c)
		sort.Ints(m.cals)
	} else {
		if c > m.cals[0] {
			m.cals[0] = c
			sort.Ints(m.cals)
		}
	}
}

func (m *sortedgroup) sum() (sum int) {
	for _, c := range m.cals {
		sum += c
	}
	return sum
}

func calsummer(input string, groupsize int) (groupcal int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var partsum int
	mx := &sortedgroup{groupsize: groupsize}
	for scanner.Scan() {
		line := scanner.Text()
		c, err := strconv.Atoi(line)
		if line != "" && err == nil {
			partsum += c
		} else {
			mx.add(partsum)
			partsum = 0
		}
	}
	mx.add(partsum)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return mx.sum()
}

func main() {
	input := "input.txt"

	fmt.Println("Task 1 - # calories of top elve \t =  ", calsummer(input, 1))
	fmt.Println("Task 2 - # cal sum of top 3 elves \t =  ", calsummer(input, 3))
}
