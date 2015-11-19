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
	return g.verticalOf(o) && g.Y == o.Y || g.horizontalOf(o) && g.X == o.X
}

func (g *GridPoint) verticalOf(o GridPoint) bool {
	return (g.X == o.X-1 || g.X == o.X+1)
}

func (g *GridPoint) horizontalOf(o GridPoint) bool {
	return (g.Y == o.Y-1 || g.Y == o.Y+1)
}

func (g *GridPoint) verticalDistance(o GridPoint) int {

	f := math.Abs(float64(g.X) - float64(o.X))
	return int(f)
}
