package main

import "math"

func (p *point) ManhattanDistance(to *point) float64 {
	return float64(abs(p.x-to.x) + abs(p.y-to.y))
}

func (p *point) EuklidianDistance(to *point) float64 {
	x := p.x - to.x
	y := p.y - to.y
	return math.Sqrt(float64(x*x + y*y))
}

func (p *point) HeightDistance(to *point) float64 {
	var x, endx int
	dx := to.x - p.x
	if dx > 0 {
		x = p.x
		endx = to.x
		dx = 1
	} else if dx < 0 {
		x = to.x
		endx = p.x
		dx = 1
	}
	var y, endy int
	dy := to.y - p.y
	if dy > 0 {
		y = p.y
		endy = to.y
		dy = 1
	} else if dy < 0 {
		y = to.y
		endy = p.y
		dy = 1
	}
	var slope int
	for {
		moved := false
		if x != endx {
			x += dx
			moved = true
		}
		if y != endy {
			y += dy
			moved = true
		}
		if x > 0 && x < p.board.dimx && y >= 0 && y < p.board.dimy {
			slope += int(p.board.points[y][x].char)
		}
		if !moved {
			break
		}
	}
	return float64(slope)
}
