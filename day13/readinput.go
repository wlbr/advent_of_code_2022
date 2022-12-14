package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
)


func safeAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error parsing '%s' as int:  %s", s, err)
	}
	return i
}

func readInput(fname string) (left, right *packet) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()



	scanner := bufio.NewScanner(f)
	elems:=NewStack[*element]()

	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			scanner.Scan()
			line = scanner.Text()
		}
		left := &packet{}
		currentelem := &element{}

		for _, r := range line {
			switch r {
			case '[':
				elems.Push(currentelem)
				newelem:=&element{}
				currentelem.packet=append(currentelem.packet,newelem)
				currentelem = newelem
			case ']':
				currentelem = elems.Pop()
			case ','," ":
				continue
			default:
				currentelem.val += safeAtoi(r)
			}
		}

		for _, r := range line {
			switch r {
			case '[':
				elems.Push(currentelem)
				newelem:=&element{}
				currentelem.packet=append(currentelem.packet,newelem)
				currentelem = newelem
				if right.elems == nil {
					right.elems = []*element{currentelem}
				}
			case ']':
				currentelem = elems.Pop()
			case ','," ":
				continue
			default:
				currentelem.val += safeAtoi(r)
			}
		}


	return monkeys
}
