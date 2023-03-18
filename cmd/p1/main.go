package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}

type Point struct {
	X int
	Y int
}

func navigate(line string) Point {
	parts := strings.Split(line, " ")
	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("this should not happen")
	}
	if parts[0] == "forward" {
		return Point{
			X: amount,
			Y: 0,
		}
	} else if parts[0] == "up" {
		return Point{
			X: 0,
			Y: -amount,
		}
	}

	return Point{
		X: 0,
		Y: amount,
	}

}

func main() {
	lines := strings.Split(getInput(), "\n")
	pos := Point{0, 0}
	for _, line := range lines {
		amount := navigate(line)
		pos.X += amount.X
		pos.Y += amount.Y
	}
	fmt.Print("point: %+v", pos)
}
