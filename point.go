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

func (g *GridPoint) SameCoordinatesWith(o GridPoint) bool {
	return g.X == o.X && g.Y == o.Y
}

func (g *GridPoint) NeighborOf(o GridPoint) bool {
	if g.SameCoordinatesWith(o) {
		return false
	}
	return g.verticalDistance(o) == 1 && g.Y == o.Y || g.horizontalDistance(o) == 1 && g.X == o.X
}

func (g *GridPoint) verticalDistance(o GridPoint) int {
	return distance(g.X, o.X)
}

func (g *GridPoint) horizontalDistance(o GridPoint) int {
	return distance(g.Y, o.Y)
}

func distance(n, m int) int {
	f := math.Abs(float64(n) - float64(m))
	return int(f)
}
