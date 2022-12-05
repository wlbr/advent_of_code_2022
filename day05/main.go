package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseIndex(lines []string) (stacks []*stack) {
	numline := lines[len(lines)-1]
	for i := 0; i < len(numline); i = i + 4 {
		m := i + 3
		if m > len(numline) {
			m = len(numline)
		}
		// nstr := strings.Trim(lines[len(lines)-1][i:m], " ")
		// n, err := strconv.Atoi(nstr)
		// if err != nil {
		// 	log.Fatalf("Error parsing number '%s':  %s", nstr, err)
		// }
		stacks = append(stacks, &stack{})
	}
	return stacks
}

func parseStacks(lines []string, stacks []*stack) []*stack {
	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		var s int
		for j := 0; j < len(line); j = j + 4 {
			m := j + 1
			if m > len(line) {
				m = len(line)
			}
			if line[m] != ' ' {
				stacks[s].push(rune(line[m]))
			}
			s++
		}
	}
	return stacks
}

type command struct {
	amount int
	from   int
	to     int
}

func parseCommands(lines []string) (commands []*command) {
	for _, line := range lines {
		c := &command{}
		fmt.Sscanf(line, "move %d from %d to %d", &c.amount, &c.from, &c.to)
		c.from--
		c.to--
		commands = append(commands, c)
	}

	return commands
}

func parseInput(input string) (stacks []*stack, commands []*command) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			lines = append(lines, line)
		}
	}
	stacks = parseIndex(lines)
	stacks = parseStacks(lines, stacks)

	lines = []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	commands = parseCommands(lines)
	return stacks, commands
}

func task1(input string) (tops string) {
	stacks, commands := parseInput(input)

	for _, c := range commands {
		for i := 0; i < c.amount; i++ {
			stacks[c.to].push(stacks[c.from].pop())
		}
	}
	for _, s := range stacks {
		tops += string(s.pop())
	}
	return tops
}

func task2(input string) (score string) {
	stacks, commands := parseInput(input)

	for _, c := range commands {
		stacks[c.to].content = stacks[c.to].content + stacks[c.from].content[stacks[c.from].len()-c.amount:stacks[c.from].len()]
		stacks[c.from].content = stacks[c.from].content[:stacks[c.from].len()-c.amount]

	}
	for _, s := range stacks {
		score += string(s.pop())
	}
	return score
}

func main() {
	input := "input.txt"

	fmt.Println("Task 1 - CrateMover 9000 - top of stacks \t =  ", task1(input))
	fmt.Println("Task 2 - CrateMover 9001 - top of stacks \t =  ", task2(input))
}
