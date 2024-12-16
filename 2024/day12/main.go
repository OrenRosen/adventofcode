package main

import (
	"fmt"
	"time"

	"github.com/OrenRosen/adventofcode/2024/textutil"
)

const (
	UP    = 0
	DOWN  = 1
	LEFT  = 2
	RIGHT = 3
)

func main() {
	defer func(t time.Time) {
		fmt.Println("time", time.Since(t))
	}(time.Now())

	file := "./day12/input.txt"
	matrix, err := textutil.GetMatrix(file)
	if err != nil {
		panic(err)
	}

	part1(matrix)
	part2(matrix)
}

func part1(matrix [][]string) {
	defer func(t time.Time) {
		fmt.Println("part 1 time", time.Since(t))
	}(time.Now())

	total := 0
	for x := range matrix {
		for y := range matrix[x] {
			if WasVisited(x, y) {
				continue
			}

			s := service{
				matrix: matrix,
			}

			BFSMatrix(matrix, x, y, s.ShouldTravel, s.HandleVisit, s.HandleTravelBlockedPart1)
			total += s.area * s.parameter
		}
	}

	fmt.Println("------------------ part 1:", total)
}

func part2(matrix [][]string) {
	defer func(t time.Time) {
		fmt.Println("part 2 time", time.Since(t))
	}(time.Now())

	ResetVisited()

	total := 0
	for x := range matrix {
		for y := range matrix[x] {
			if WasVisited(x, y) {
				continue
			}

			s := service{
				matrix: matrix,
			}

			BFSMatrix(matrix, x, y, s.ShouldTravel, s.HandleVisit, s.HandleTravelBlockedPart2)
			total += s.area * s.parameter
		}
	}

	fmt.Println("------------------ total", total)
}

type service struct {
	matrix [][]string

	area      int
	parameter int

	// only used in part2
	preventedTiles map[Point][]bool
}

func (s *service) AddToPreventedTiles(to Point, direction int) {
	if s.preventedTiles == nil {
		s.preventedTiles = make(map[Point][]bool)
	}

	if s.preventedTiles == nil {
		s.preventedTiles = make(map[Point][]bool)
	}

	if s.preventedTiles[to] == nil {
		s.preventedTiles[to] = make([]bool, 4)
	}

	s.preventedTiles[to][direction] = true
}

func (s *service) ShouldTravel(from, to Point) bool {
	fromVal := s.matrix[from.X][from.Y]
	toVal := s.matrix[to.X][to.Y]

	return fromVal == toVal
}

func (s *service) HandleVisit(p Point) {
	s.area++
}

func (s *service) HandleTravelBlockedPart1(from, to Point, becauseAlreadyVisited bool) {
	if becauseAlreadyVisited {
		// we need to add to perimeter, only if it's not the same value
		fromValue := s.matrix[from.X][from.Y]
		toValue := s.matrix[to.X][to.Y]
		if fromValue == toValue {
			return
		}
	}

	s.parameter++
}

func (s *service) HandleTravelBlockedPart2(from, to Point, becauseAlreadyVisited bool) {
	direction := getDirection(from, to)
	toValue, fromValue := "", ""

	if IsOutOfBoundsMatrix(s.matrix, to) {
		s.AddToPreventedTiles(to, direction)
	} else {
		toValue = s.matrix[to.X][to.Y]
		fromValue = s.matrix[from.X][from.Y]

		// only add to prevented if it's not the same value
		if fromValue != toValue {
			s.AddToPreventedTiles(to, direction)
		}
	}

	// if we are blocked because it was already visited, we need to add to perimeter, only if it's not the same value
	if becauseAlreadyVisited && fromValue == toValue {
		return
	}

	adjustentTiles := []Point{}
	isPreventedFromRightOrLeft := to.X == from.X
	if isPreventedFromRightOrLeft {
		adjustentTiles = []Point{
			{X: to.X + 1, Y: to.Y},
			{X: to.X - 1, Y: to.Y},
		}
	} else {
		adjustentTiles = []Point{
			{X: to.X, Y: to.Y + 1},
			{X: to.X, Y: to.Y - 1},
		}
	}

	for _, t := range adjustentTiles {
		preventedDirections, ok := s.preventedTiles[t]
		if ok && preventedDirections[direction] {
			return
		}
	}

	s.parameter++
}

func getDirection(from, to Point) int {
	if from.X < to.X {
		return DOWN
	} else if from.X > to.X {
		return UP
	} else if from.Y < to.Y {
		return RIGHT
	} else if from.Y > to.Y {
		return LEFT
	}

	panic("Invalid direction")
}

func IsOutOfBoundsMatrix(matrix [][]string, p Point) bool {
	return p.X < 0 || p.Y < 0 || p.X >= len(matrix) || p.Y >= len(matrix[p.X])
}
