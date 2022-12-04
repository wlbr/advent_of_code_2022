package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func odd(x int) bool {
	return x%2 == 1
}

func charValue(c rune) int {
	offsetUppercase := int('A') + int('a') - int('z') - 1
	offsetLowercase := int('a')

	if unicode.IsUpper(c) {
		return int(c) - offsetUppercase + 1
	}
	return int(c) - offsetLowercase + 1
}

func task1(input string) (score int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if odd(len(line)) {
			log.Fatalf("Backpack is not balanced. #Parcels=%d ", len(line))
		}
		doubles := make([]int, 53)
		checklist := make([]int, 53)
		for _, c := range line[:len(line)/2] {
			checklist[charValue(c)]++
		}
		for _, c := range line[len(line)/2:] {
			if checklist[charValue(c)] > 0 {
				doubles[charValue(c)]++
			}
		}
		for i, d := range doubles {
			if d > 0 {
				score += i
			}
		}
	}

	return score
}

func task2(input string) (score int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var badges [][]int
	var groupcount int
	for scanner.Scan() {

		line := scanner.Text()
		if odd(len(line)) {
			log.Fatalf("Backpack is not balanced. #Parcels=%d ", len(line))
		}
		if groupcount == 0 {
			badges = make([][]int, 3)
			for i, _ := range badges {
				badges[i] = make([]int, 53)
			}
		}
		for _, c := range line {
			badges[groupcount][charValue(c)]++
		}
		groupcount++
		if groupcount == 3 {
			groupcount = 0

			for i := 0; i < 53; i++ {
				isbadge := true
				for g := 0; g < 3; g++ {
					if badges[g][i] == 0 {
						isbadge = false
					}
				}
				if isbadge {
					score += i
				}
			}
		}
	}

	return score
}

func main() {
	input := "input.txt"

	fmt.Println("Task 1 - # sum of priorities of double items \t =  ", task1(input))
	fmt.Println("Task 2 - # sum of priorities of badges       \t =  ", task2(input))
}
