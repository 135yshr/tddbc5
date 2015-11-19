package tddbc5

import (
	"fmt"
)

type GridPoint struct {
	X int
	Y int
}

func (g *GridPoint) Notation() string {
	return fmt.Sprintf("(%d,%d)", g.X, g.Y)
}

func (g *GridPoint) SameCoordinatesWith(other GridPoint) bool {
	return g.X == other.X && g.Y == other.Y
}

func (g *GridPoint) NeighborOf(other GridPoint) bool {
	if g.SameCoordinatesWith(other) {
		return false
	}
	return (g.X == other.X-1 || g.X == other.X+1) && g.Y == other.Y || (g.Y == other.Y-1 || g.Y == other.Y+1) && g.X == other.X
}

func (g *GridPoint) VerticalOf(o GridPoint) bool {
	return (g.X == o.X-1 || g.X == o.X+1)
}
