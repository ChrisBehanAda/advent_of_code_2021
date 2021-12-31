package day17

import "fmt"

type rectangle struct {
	x1, x2 int
	y1, y2 int
}

type point struct {
	x, y int
}

func Solve() {
	targetArea := rectangle{x1: 150, x2: 171, y1: -70, y2: -129}
	yVel := highestYVel(targetArea)
	fmt.Printf("Highest y position is: %v\n", yVel)
	posibilities := viableVelocities(targetArea)
	fmt.Printf("%v initial velocities that can hit the target area.\n", posibilities)
}

func path(xVel, yVel int, target rectangle) ([]point, bool) {
	x, y := 0, 0
	path := []point{}
	hit := false
	for x <= target.x2 && y >= target.y2 {
		x += xVel
		y += yVel
		if xVel > 0 {
			xVel--
		}
		yVel--
		path = append(path, point{x: x, y: y})
		if !hit {
			hit = collision(x, y, target)
		}
	}
	return path, hit
}

func collision(x, y int, target rectangle) bool {
	if x >= target.x1 && x <= target.x2 && y <= target.y1 && y >= target.y2 {
		return true
	}
	return false
}

func highestYVel(target rectangle) int {
	hiY := -1
	for x := 0; x < abs(target.x2); x++ {
		for y := 0; y < abs(target.y2); y++ {
			path, hit := path(x, y, target)
			if hit && highestYPos(path) > hiY {
				hiY = highestYPos(path)
			}
		}
	}
	return hiY
}

func viableVelocities(target rectangle) int {
	count := 0
	for x := 0; x <= abs(target.x2); x++ {
		for y := -abs(target.y2); y <= abs(target.y2); y++ {
			_, hit := path(x, y, target)
			if hit {
				count++
			}
		}
	}
	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func highestYPos(path []point) int {
	y := 0
	for _, p := range path {
		if p.y > y {
			y = p.y
		}
	}
	return y
}
