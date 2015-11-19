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

func (g *GridPoint) HasSameCoordinatesWith(a GridPoint) bool {
	return g.X == a.X && g.Y == a.Y
}

func (g *GridPoint) IsNeighborOf(a GridPoint) bool {
	if g.horizontalDistance(a) == 1 && g.verticalDistance(a) == 0 {
		return true
	}
	if g.horizontalDistance(a) == 0 && g.verticalDistance(a) == 1 {
		return true
	}
	return false
}

func (g *GridPoint) verticalDistance(a GridPoint) int {
	return distance(g.Y, a.Y)
}

func (g *GridPoint) horizontalDistance(a GridPoint) int {
	return distance(g.X, a.X)
}

func distance(f, t int) int {
	r := math.Abs(float64(f) - float64(t))
	return int(r)
}
