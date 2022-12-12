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
	char  rune
	board *board
}

func NewPoint(x, y int, board *board, char rune) *point {
	p := point{x: x, y: y, board: board, char: char}
	switch p.char {
	case 'S':
		p.char = 'a'
	case 'E':
		p.char = 'z'
	}
	return &p
}

func (p *point) String() string {
	return fmt.Sprintf("[%d,%d]=%c", p.x, p.y, p.char)
}

func (p *point) checkAndAddPointToAdjacents(other *point, adj *[]astar.Pather) {
	if other.char-p.char <= 1 {
		*adj = append(*adj, other)
	}
}

func (p *point) PathNeighbors() []astar.Pather {
	var adj []astar.Pather

	if p.x > 0 {
		p.checkAndAddPointToAdjacents(p.board.points[p.y][p.x-1], &adj)
	}
	if p.x < p.board.dimx-1 {
		p.checkAndAddPointToAdjacents(p.board.points[p.y][p.x+1], &adj)
	}
	if p.y > 0 {
		p.checkAndAddPointToAdjacents(p.board.points[p.y-1][p.x], &adj)
	}
	if p.y < p.board.dimy-1 {
		p.checkAndAddPointToAdjacents(p.board.points[p.y+1][p.x], &adj)
	}

	return adj
}

func (p *point) PathNeighborCost(to astar.Pather) float64 {
	top := to.(*point)
	return float64(abs(int(p.char - top.char)))
}

func (p *point) PathEstimatedCost(to astar.Pather) float64 {
	//c := p.HeightDistance(to.(*point))
	//c := p.EuklidianDistance(to.(*point))
	c := p.ManhattanDistance(to.(*point))
	return c
}

type board struct {
	fields []string
	points [][]*point
	start  *point
	target *point
	dimx   int
	dimy   int
}

func task1(fname string) int {
	b := readInput(fname)

	path, _, found := astar.Path(b.start, b.target)
	if !found {
		fmt.Println("No path found")
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

	fmt.Printf("Task 1 - steps to target    \t: %d \n", t1)
	fmt.Printf("Task 2 -                    \t: %d \n\n", t2)

	fmt.Println("Time task 1: ", afterTask1.Sub(startOverall))
	fmt.Println("Time task 2: ", afterTask2.Sub(afterTask1))
	fmt.Println("Total time: ", afterTask2.Sub(startOverall))
}
