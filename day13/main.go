package main

import (
	"fmt"
	"time"
)

type element struct {
	values element
	packet *packet
}

type packet struct {
	elems []*element
}

func task1(fname string) int {
	left, right := readInput(fname)
	return 0
}

func task2(fname string) int {
	//monkeys := readInput(fname)
	return 1
}

func main() {
	startOverall := time.Now()
	input := "input.txt"
	t1 := task1(input)
	afterTask1 := time.Now()
	t2 := task2(input)
	afterTask2 := time.Now()

	fmt.Printf("Task 1 - after 20 round    \t:  %d \n", t1)
	fmt.Printf("Task 2 - after 10000 rounds \t: %d \n\n", t2)

	fmt.Println("Time task 1: ", afterTask1.Sub(startOverall))
	fmt.Println("Time task 2: ", afterTask2.Sub(afterTask1))
	fmt.Println("Total time: ", afterTask2.Sub(startOverall))
}
