package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(fname string) (buffer string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		buffer += line
	}
	return buffer
}

func findsegment(input string, markersize int) int {
	buffer := readInput(input)
	markerset := ""
	for pos, c := range buffer {
		if p := strings.Index(markerset, string(c)); p != -1 {
			markerset += string(c)
			markerset = markerset[p+1:]
		} else {
			markerset += string(c)
			if len(markerset) == markersize {
				return pos + 1
			}
		}

	}
	return 0
}

func task2(input string) int {
	return 0
}

func main() {
	input := "input.txt"

	fmt.Println("Task 1 - start marker at pos   \t =  ", findsegment(input, 4))
	fmt.Println("Task 2 - message marker at pos \t =  ", findsegment(input, 14))
}
