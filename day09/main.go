package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sign(a int) int {
	if a < 0 {
		return -1
	}
	return 1
}

type direction int

func (d direction) String() string {
	switch d {
	case NORTH:
		return "U"
	case SOUTH:
		return "D"
	case EAST:
		return "R"
	case WEST:
		return "L"
	}
	return "UNKNOWN"
}

const (
	NORTH direction = iota
	SOUTH
	EAST
	WEST
)

type motion struct {
	dir      direction
	distance int
}

func (m motion) String() string {
	return fmt.Sprintf("%s %d", m.dir, m.distance)
}

func readInput(fname string) (motions []*motion) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	var cmds string
	var dis int
	for _, end := fmt.Fscanf(f, "%s %d", &cmds, &dis); end != io.EOF; _, end = fmt.Fscanf(f, "%s %d", &cmds, &dis) {
		var cmd direction
		switch cmds {
		case "U":
			cmd = NORTH
		case "D":
			cmd = SOUTH
		case "R":
			cmd = EAST
		case "L":
			cmd = WEST
		default:
			log.Fatalf("Unknown motion '%s'", cmds)
		}
		motions = append(motions, &motion{cmd, dis})
	}

	return motions
}

type point struct {
	x, y int
}

func (p point) String() string {
	return fmt.Sprintf("[%d,%d]", p.x, p.y)
}

func follow(currentTail, currentHead point) (newTail point) {
	xdistance := currentHead.x - currentTail.x
	ydistance := currentHead.y - currentTail.y

	newTail = point{currentTail.x, currentTail.y}

	if (abs(xdistance) > 1 && abs(ydistance) > 0) || (abs(ydistance) > 1 && abs(xdistance) > 0) {
		newTail.x = currentTail.x + sign(xdistance)*min(abs(xdistance), 1)
		newTail.y = currentTail.y + sign(ydistance)*min(abs(ydistance), 1)
	} else {
		if abs(xdistance) > 1 {
			step := sign(xdistance) * min(abs(xdistance), 1)
			newTail.x = currentTail.x + step
		}
		if abs(ydistance) > 1 {
			step := sign(ydistance) * min(abs(ydistance), 1)
			newTail.y = currentTail.y + step
		}
	}

	return newTail
}

func moveHead(rope []point, cmd *motion) {
	switch cmd.dir {
	case NORTH:
		rope[0].y++
	case SOUTH:
		rope[0].y--
	case EAST:
		rope[0].x++
	case WEST:
		rope[0].x--
	}
}

func walkRope(input string, numknots int) int {
	motions := readInput(input)
	visited := make(map[string]bool)
	rope := make([]point, numknots)
	for i := 0; i < numknots; i++ {
		rope[i] = point{0, 0}
	}

	for _, cmd := range motions {
		for i := 0; i < cmd.distance; i++ {
			moveHead(rope, cmd)
			for k := 1; k < numknots; k++ {
				rope[k] = follow(rope[k], rope[k-1])
			}
			visited[rope[len(rope)-1].String()] = true
		}
	}
	return len(visited)
}

func main() {
	startOverall := time.Now()
	input := "input.txt"
	afterParsing := time.Now()

	t1 := walkRope(input, 2)
	afterTask1 := time.Now()
	t2 := walkRope(input, 10)
	afterTask2 := time.Now()

	fmt.Println("Task 1 - positions visited by the tail of 2   =  ", t1)
	fmt.Println("Task 2 - positions visited by the tail of 10  =  ", t2)

	fmt.Println("\nTime parsing input: ", afterParsing.Sub(startOverall))
	fmt.Println("Time task 1: ", afterTask1.Sub(afterParsing))
	fmt.Println("Time task 2: ", afterTask2.Sub(afterTask1))
	fmt.Println("Total time: ", afterTask2.Sub(startOverall))
}
