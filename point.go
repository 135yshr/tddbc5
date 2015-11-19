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
	return g.SameCoordinatesWith(other) == false
}
