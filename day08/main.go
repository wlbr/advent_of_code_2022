package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

//type board [][]int

func readInput(fname string) (buffer [][]int) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var intline []int
		for _, s := range line {
			i, err := strconv.Atoi(string(s))
			if err != nil {
				log.Fatalf("Error parsing '%c':  %s", s, err)
			}
			intline = append(intline, i)
		}
		buffer = append(buffer, intline)
	}
	return buffer
}

type point [2]int

func (p point) x() int {
	return p[0]
}

func (p point) y() int {
	return p[1]
}

func (p point) move(dir direction) (newpoint point) {
	newpoint[0] = p[0] + dir[0]
	newpoint[1] = p[1] + dir[1]
	return newpoint
}

type direction [2]int

func (d direction) x() int {
	return d[0]
}

func (d direction) y() int {
	return d[1]
}

func walk(b [][]int, start point, d direction, collector map[string]int) {
	x := start.x()
	y := start.y()
	last := -1
	for y <= len(b)-1 && y >= 0 && x <= len(b[0])-1 && x >= 0 {

		p := b[y][x]
		if p > last {
			key := fmt.Sprintf("%d,%d", x, y)
			collector[key]++
			last = p
		}

		x += d.x()
		y += d.y()

	}
}

func checkPoint(board [][]int, p point, d []direction, collector map[string]int) {
	var x, y int

	for _, d := range d {
		switch d[0] {
		case -1:
			x = len(board[0]) - 1
		default:
			x = p.x()
		}
		switch d[1] {
		case -1:
			y = len(board) - 1
		default:
			y = p.y()
		}
		walk(board, point{x, y}, d, collector)
	}
}

func task1(input string) (counttrees int) {
	board := readInput(input)
	collector := make(map[string]int)
	for x := 0; x < len(board[0]); x++ {
		checkPoint(board, point{x, 0}, []direction{{0, 1}, {0, -1}}, collector)
	}

	for y := 0; y < len(board); y++ {
		checkPoint(board, point{0, y}, []direction{{1, 0}, {-1, 0}}, collector)
	}

	return len(collector)
}

func viewdistance(b [][]int, start point, d direction, collector map[string]int) int {
	x := start.x()
	y := start.y()
	distance := 0
	currentheight := b[y][x]
	for y < len(b)-1 && y > 0 && x < len(b[0])-1 && x > 0 {
		x += d.x()
		y += d.y()
		p := b[y][x]
		if p >= currentheight {
			return distance + 1
		}
		distance++

	}
	return distance
}

func scenicScore(board [][]int, p point, d []direction, collector map[string]int) int {
	score := 1
	var distances []int
	for _, d := range d {
		dis := viewdistance(board, p, d, collector)
		distances = append(distances, dis)
	}
	for _, distance := range distances {
		score = score * distance
	}
	return score
}

func task2(input string) (max int) {
	board := readInput(input)
	collector := make(map[string]int)
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			sscore := scenicScore(board, point{x, y}, []direction{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}, collector)
			if sscore > max {
				max = sscore
			}
		}
	}

	return max
}

func main() {
	startOverall := time.Now()
	input := "input.txt"
	afterParsing := time.Now()

	t1 := task1(input)
	afterTask1 := time.Now()
	t2 := task2(input)
	afterTask2 := time.Now()
	fmt.Println("Task 1 - visible trees        \t =  ", t1)
	fmt.Println("Task 2 - highest scenic score \t =  ", t2)

	fmt.Println("\nTime parsing input: ", afterParsing.Sub(startOverall))
	fmt.Println("Time task 1: ", afterTask1.Sub(afterParsing))
	fmt.Println("Time task 2: ", afterTask2.Sub(afterTask1))
	fmt.Println("Total time: ", afterTask2.Sub(startOverall))

}
