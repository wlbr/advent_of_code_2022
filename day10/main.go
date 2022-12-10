package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type instruction struct {
	cmd      string
	argument int
}

func (i instruction) String() string {
	return fmt.Sprintf("%s %d", i.cmd, i.argument)
}

func readInput(fname string) (cmds []*instruction) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	var cmd string
	var arg int
	for _, end := fmt.Fscanf(f, "%s %d", &cmd, &arg); end != io.EOF; _, end = fmt.Fscanf(f, "%s %d", &cmd, &arg) {

		cmds = append(cmds, &instruction{cmd, arg})
	}

	return cmds
}
func getSignalStrength(signalstrength, cycle, x int, cmd *instruction) int {
	if (cycle+20)%40 == 0 {
		signalstrength += cycle * x
	}
	return signalstrength
}

func task1(fname string) (signalstrength int) {
	cmds := readInput(fname)

	cycle := 1
	x := 1
	for _, cmd := range cmds {
		switch cmd.cmd {
		case "noop":
			signalstrength = getSignalStrength(signalstrength, cycle, x, cmd)
			cycle++
		case "addx":
			for i := 0; i < 2; i++ {
				signalstrength = getSignalStrength(signalstrength, cycle, x, cmd)
				cycle++
			}
			x += cmd.argument
		}
	}

	return signalstrength
}

type crt struct {
	xdim, ydim int
	pixels     []int
}

func NewCrt(xdim, ydim int) *crt {

	return &crt{pixels: make([]int, xdim*ydim), xdim: xdim, ydim: ydim}
}

func (c crt) String() string {
	var str string
	for i, p := range c.pixels {
		if i%c.xdim == 0 {
			str += "\n"
		}
		if p == 0 {
			str += "."
		} else {
			str += "#"
		}
	}

	return str
}

func (c crt) draw(x int) {
	if x < len(c.pixels) && x >= 0 {
		c.pixels[x]++
	} else {
		log.Fatal("x out of range", x)
	}

}

func (c crt) drawXY(x, y int) {
	if x+y*c.xdim < len(c.pixels) && x+y*c.xdim >= 0 {
		c.pixels[x+y*c.xdim] += 1
	} else {
		log.Fatal("x or y out of range", x)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (c crt) checkSpriteCoverage(spritecenter, cycle int) bool {
	if cycle > c.xdim-1 || cycle < 0 {
		log.Fatalf("point our of range")
	}
	start := max(0, spritecenter-1)
	end := min(c.xdim-1, spritecenter+1)
	for i := start; i <= end; i++ {
		if cycle >= start && cycle <= end {
			return true
		}
	}
	return false
}

func (c crt) incCoords(x, y int) (int, int) {
	x++
	if x%c.xdim == 0 {
		y++
		x = 0
	}
	return x, y
}

func task2(fname string) string {
	cmds := readInput(fname)
	crt := NewCrt(40, 6)
	spritecenter := 1
	y := 0
	x := 0
	for _, cmd := range cmds {
		switch cmd.cmd {
		case "noop":
			if crt.checkSpriteCoverage(spritecenter, x) {
				crt.drawXY(x, y)
			}
			x, y = crt.incCoords(x, y)
		case "addx":
			for i := 0; i < 2; i++ {
				if crt.checkSpriteCoverage(spritecenter, x) {
					crt.drawXY(x, y)
				}
				x, y = crt.incCoords(x, y)
			}
			spritecenter += cmd.argument
		}
	}

	return crt.String()
}

func main() {
	startOverall := time.Now()
	input := "input.txt"
	t1 := task1(input)
	afterTask1 := time.Now()
	t2 := task2(input)
	afterTask2 := time.Now()

	fmt.Println("Task 1 - positions visited by the tail of 2   =  ", t1)
	fmt.Printf("Task 2 - chars are: %s \n\n", t2)

	fmt.Println("Time task 1: ", afterTask1.Sub(startOverall))
	fmt.Println("Time task 2: ", afterTask2.Sub(afterTask1))
	fmt.Println("Total time: ", afterTask2.Sub(startOverall))
}
