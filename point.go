package tddbc5

type Point struct {
	X int
	Y int
}

func (p *Point) IsNeighbor(p2 Point) bool {
	return true
}
