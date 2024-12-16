package main

import (
	"container/list"
)

type Point struct {
	X, Y int
}

func (p Point) IsAdjacent(other Point) bool {
	return (p.X == other.X && (p.Y == other.Y+1 || p.Y == other.Y-1)) || (p.Y == other.Y && (p.X == other.X+1 || p.X == other.X-1))
}

type ShouldTravelFunc func(from, to Point) bool
type HandleVisitFunc func(p Point)
type HandleTravelBlockedFunc func(from, to Point, becauseAlreadyVisited bool)

var visited = map[Point]bool{}

func BFSMatrix[T comparable](matrix [][]T, startX, startY int, shouldTravel ShouldTravelFunc, handleVisitFunc HandleVisitFunc, handleTravelBlockedFunc HandleTravelBlockedFunc) {
	rows := len(matrix)
	cols := len(matrix[0])

	// Directions for moving in 4 possible ways (up, down, left, right)
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Queue for BFS
	queue := list.New()
	queue.PushBack([2]int{startX, startY})

	p := Point{startX, startY}
	visited[p] = true

	for queue.Len() > 0 {
		// Dequeue the front cell
		current := queue.Remove(queue.Front()).([2]int)
		x, y := current[0], current[1]
		currPoint := Point{x, y}
		handleVisitFunc(currPoint)

		// Explore neighbors
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			nextPoint := Point{newX, newY}

			if visited[nextPoint] {
				handleTravelBlockedFunc(currPoint, nextPoint, true)
				continue
			}

			if newX < 0 || newX >= rows || newY < 0 || newY >= cols {
				handleTravelBlockedFunc(currPoint, nextPoint, false)
				continue
			}

			if !shouldTravel(currPoint, nextPoint) {
				handleTravelBlockedFunc(currPoint, nextPoint, false)
				continue
			}

			queue.PushBack([2]int{newX, newY})
			visited[nextPoint] = true
		}
	}
}

func WasVisited(x, y int) bool {
	p := Point{x, y}
	return visited[p]
}

func ResetVisited() {
	visited = map[Point]bool{}
}
