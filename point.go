package tddbc5

type Point struct {
	X int
	Y int
}

func (p *Point) IsNeighbor(p2 Point) bool {
	if p.X != p2.X && p.Y != p2.Y {
		return false
	}
	return true
}
