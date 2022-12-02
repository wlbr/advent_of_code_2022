package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func codeToValue(code string) (value int) {
	switch code {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	}
	return 0
}

const (
	win  = 6
	loss = 0
	draw = 3
)

func Score(left, right string) (score int) {
	if codeToValue(left) == codeToValue(right) {
		score = draw
	}
	switch right {
	case "X":
		switch left {
		case "B":
			score = loss
		case "C":
			score = win
		}
	case "Y":
		switch left {
		case "A":
			score = win
		case "C":
			score = loss
		}
	case "Z":
		switch left {
		case "A":
			score = loss
		case "B":
			score = win
		}
	}
	return score + codeToValue(right)
}

func task1(input string) (score int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()
	fmt.Println()

	var left, right string
	for _, end := fmt.Fscanln(f, &left, &right); end != io.EOF; _, end = fmt.Fscanln(f, &left, &right) {
		score += Score(left, right)
	}

	return score
}

func makeMove(opponent, goal string) string {
	switch goal {
	case "Y": //draw:
		switch opponent {
		case "A":
			return "X"
		case "B":
			return "Y"
		case "C":
			return "Z"
		}
	case "Z": //win:
		switch opponent {
		case "A":
			return "Y"
		case "B":
			return "Z"
		case "C":
			return "X"
		}
	case "X": //loss:
		switch opponent {
		case "A":
			return "Z"
		case "B":
			return "X"
		case "C":
			return "Y"
		}
	}
	return ""
}

func task2(input string) (score int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()
	fmt.Println()

	var left, res string
	for _, end := fmt.Fscanln(f, &left, &res); end != io.EOF; _, end = fmt.Fscanln(f, &left, &res) {
		right := makeMove(left, res)
		score += Score(left, right)
	}

	return score
}

func main() {
	input := "input.txt"

	fmt.Println("Task 1 - # calories of top elve \t =  ", task1(input))
	fmt.Println("Task 2 - # cal sum of top 3 elves \t =  ", task2(input))
}
