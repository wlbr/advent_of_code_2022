package main

import (
	"bufio"
	"log"
	"os"
)

func readInput(fname string) *board {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	board := &board{}
	var y int
	for scanner.Scan() {
		line := scanner.Text()
		var pline []*point
		board.fields = append(board.fields, line)
		for x, r := range line {
			p := NewPoint(x, y, board, r)
			pline = append(pline, p)
			switch r {
			case 'S':
				board.start = p
			case 'E':
				board.target = p
			}

		}
		board.points = append(board.points, pline)
		y++
	}
	board.dimx = len(board.fields[0])
	board.dimy = len(board.fields)

	return board
}
