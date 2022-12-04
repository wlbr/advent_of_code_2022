package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Assignment struct {
	start, end int
}

func (a Assignment) Contains(b Assignment) bool {
	return a.start <= b.start && a.end >= b.end
}

func (a Assignment) AnyUnneeded(b Assignment) bool {
	return a.Contains(b) || b.Contains(a)
}

func (a Assignment) Overlaps(b Assignment) bool {
	return b.start >= a.start && b.start <= a.end ||
		b.end >= a.start && b.end <= a.end
}

func (a Assignment) AnyOverlapping(b Assignment) bool {
	return a.Overlaps(b) || b.Overlaps(a)
}

func task1(input string) (score int) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	var left, right Assignment
	for _, end := fmt.Fscanf(f, "%d-%d,%d-%d", &left.start, &left.end, &right.start, &right.end); end != io.EOF; _, end = fmt.Fscanf(f, "%d-%d,%d-%d", &left.start, &left.end, &right.start, &right.end) {
		if left.AnyUnneeded(right) {
			score++
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

	var left, right Assignment
	for _, end := fmt.Fscanf(f, "%d-%d,%d-%d", &left.start, &left.end, &right.start, &right.end); end != io.EOF; _, end = fmt.Fscanf(f, "%d-%d,%d-%d", &left.start, &left.end, &right.start, &right.end) {
		if left.AnyOverlapping(right) {
			score++
		}
	}

	return score
}

func main() {
	input := "input.txt"

	fmt.Println("Task 1 - # contained assignments \t =  ", task1(input))
	fmt.Println("Task 2 - # overlapping assignments \t =  ", task2(input))
}
