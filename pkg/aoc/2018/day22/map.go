package day22

import (
	"strings"

	"gitlab.com/leononame/advent-of-code/pkg/geo"
)

const (
	rocky = iota
	wet
	narrow
)
const overSize = 20

var tiles = map[int]byte{
	rocky:  '.',
	wet:    '=',
	narrow: '|',
}

type Region struct {
	Location            geo.Point
	Erosion, Geological int
	Terrain             int
}

type Map struct {
	Data   [][]*Region
	Depth  int
	Target geo.Point
}

func (m *Map) get(p geo.Point) *Region {
	return m.Data[p.GetY()][p.GetX()]
}

func (m *Map) calcGeo(p geo.Point) int {
	x, y := p.GetX(), p.GetY()
	switch {
	case x == m.Target.GetX() && y == m.Target.GetY():
		return 0
	case x == 0:
		return y * 48271
	case y == 0:
		return x * 16807
	default:
		p1 := p.Left()
		p2 := p.Up()
		return m.get(p1).Erosion * m.get(p2).Erosion
	}
}

func (m *Map) calcRisk() int {
	risk := 0
	for y := 0; y < m.Target.GetY()+1; y++ {
		for x := 0; x < m.Target.GetX()+1; x++ {
			risk += m.Data[y][x].Terrain
		}
	}
	return risk
}

func (m *Map) String() string {
	var sb strings.Builder
	for y := 0; y < m.Target.GetY()+1; y++ {
		for x := 0; x < m.Target.GetX()+1; x++ {
			b := tiles[m.Data[y][x].Terrain]
			sb.WriteByte(b)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (m *Map) neighbours(p geo.Point) []*Region {
	ns := p.Neighbours()
	var rs []*Region
	for _, neighbour := range ns {
		if neighbour.GetX() < 0 ||
			neighbour.GetY() < 0 ||
			neighbour.GetX() >= m.Target.GetX()+overSize ||
			neighbour.GetY() >= m.Target.GetY()+overSize {
			continue
		}
		r := m.get(neighbour)
		rs = append(rs, r)
	}
	return rs
}
