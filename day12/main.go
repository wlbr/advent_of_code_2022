package main

import (
	"fmt"
	"time"

	astar "github.com/beefsack/go-astar"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type point struct {
	x     int
	y     int
	value int
	char  rune
	board *board
}

func NewPoint(x, y int, board *board, char rune) *point {
	p := point{x: x, y: y, board: board, char: char}
	p.value = int(char) - 'a' + 1
	switch p.char {
	case 'S':
		p.value = 0
	case 'E':
		p.value = int('z') - 'a' + 2
	}
	return &p
}

func (p *point) String() string {
	return fmt.Sprintf("[%d,%d]=%c(%d)", p.x, p.y, p.char, p.value)
}

func (p *point) ManhattanDistance(to *point) float64 {
	return float64(abs(p.x-to.x) + abs(p.y-to.y))
}

func addPointToAdjacents(p *point, adj *[]astar.Pather) {
	if p != nil && p != p.board.start {
		*adj = append(*adj, p)
	}
}

func (p *point) PathNeighbors() []astar.Pather {
	var adj []astar.Pather

	if p.x > 0 && abs(p.value-p.board.points[p.y][p.x-1].value) <= 1 {
		addPointToAdjacents(p.board.points[p.y][p.x-1], &adj)
	}
	if p.x < p.board.dimx-1 && abs(p.value-p.board.points[p.y][p.x+1].value) <= 1 {
		adj = append(adj, p.board.points[p.y][p.x+1])
	}
	if p.y > 0 && abs(p.value-p.board.points[p.y-1][p.x].value) <= 1 {
		adj = append(adj, p.board.points[p.y-1][p.x])
	}
	if p.y < p.board.dimy-1 && abs(p.value-p.board.points[p.y+1][p.x].value) <= 1 {
		adj = append(adj, p.board.points[p.y+1][p.x])
	}

	for _, a := range adj {
		fmt.Printf("%s - %s=%f ", p, a, p.ManhattanDistance(a.(*point)))
	}
	fmt.Println()
	return adj
}

func (t *point) PathNeighborCost(to astar.Pather) float64 {
	tor := to.(*point)
	return float64(tor.value)
}

func (t *point) PathEstimatedCost(to astar.Pather) float64 {
	return t.ManhattanDistance(to.(*point))
}

type board struct {
	fields []string
	points [][]*point
	start  *point
	target *point
	dimx   int
	dimy   int
}

func NewBoard() *board {
	b := board{}
	return &b
}

func (b *board) String() string {
	s := ""
	for _, line := range b.fields {
		s += line + "\b"
	}
	return s
}

func (b *board) Get(x, y int) rune {
	if x < 0 || x >= b.dimx || y < 0 || y >= b.dimy {
		return '#'
	}
	return rune(b.fields[y][x])
}

func (b *board) Val(x, y int) (value int, isTarget bool) {
	if x < 0 || x >= b.dimx || y < 0 || y >= b.dimy {
		return -1, false
	}
	if b.fields[y][x] == 'E' {
		return 0, true
	}
	return b.points[y][x].value, false
}

func task1(fname string) int {
	b := readInput(fname)

	path, _, found := astar.Path(b.start, b.target)
	if !found {
		fmt.Println("No path found")
		//	return -1
	}
	return len(path) - 1
}

func task2(fname string) int {

	return 0
}

func main() {
	startOverall := time.Now()
	input := "input.txt"
	t1 := task1(input)
	afterTask1 := time.Now()
	t2 := task2(input)
	afterTask2 := time.Now()

	fmt.Printf("Task 1 - steps to target    \t:  %d \n", t1)
	fmt.Printf("Task 2 - after 10000 rounds \t: %d \n\n", t2)

	fmt.Println("Time task 1: ", afterTask1.Sub(startOverall))
	fmt.Println("Time task 2: ", afterTask2.Sub(afterTask1))
	fmt.Println("Total time: ", afterTask2.Sub(startOverall))
}
