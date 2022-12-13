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
	return float64(abs(int(p.char-top.char))) + 1
}

func (p *point) ManhattanDistance(to *point) float64 {
	return float64(abs(p.x-to.x) + abs(p.y-to.y))
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
	b := readInput(fname)

	var startingpoints []*point
	for _, row := range b.points {
		for _, p := range row {
			if p.char == 'a' {
				startingpoints = append(startingpoints, p)
			}
		}
	}

	max := 999999999
	for _, p := range startingpoints {
		path, _, found := astar.Path(p, b.target)
		if found {
			if len(path)-1 < max {
				max = len(path) - 1
			}
		}
	}
	return max
}

func main() {
	startOverall := time.Now()
	input := "input.txt"
	t1 := task1(input)
	afterTask1 := time.Now()
	t2 := task2(input)
	afterTask2 := time.Now()

	fmt.Printf("Task 1 - steps to target    \t: %d \n", t1)
	fmt.Printf("Task 2 - shortest track     \t: %d \n\n", t2)

	fmt.Println("Time task 1: ", afterTask1.Sub(startOverall))
	fmt.Println("Time task 2: ", afterTask2.Sub(afterTask1))
	fmt.Println("Total time: ", afterTask2.Sub(startOverall))
}
