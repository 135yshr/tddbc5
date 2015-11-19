package tddbc5

import (
	"fmt"
	"math"
)

type GridPoint struct {
	X int
	Y int
}

func (g *GridPoint) Notation() string {
	return fmt.Sprintf("(%d,%d)", g.X, g.Y)
}

func (g *GridPoint) HasSameCoordinatesWith(o GridPoint) bool {
	return g.X == o.X && g.Y == o.Y
}

func (g *GridPoint) IsNeighborOf(o GridPoint) bool {
	if g.horizontalDistance(o) == 1 && g.verticalDistance(o) == 0 {
		return true
	}
	if g.horizontalDistance(o) == 0 && g.verticalDistance(o) == 1 {
		return true
	}
	return false
}

func (g *GridPoint) verticalDistance(o GridPoint) int {
	return distance(g.Y, o.Y)
}

func (g *GridPoint) horizontalDistance(o GridPoint) int {
	return distance(g.X, o.X)
}

func distance(n, m int) int {
	f := math.Abs(float64(n) - float64(m))
	return int(f)
}
