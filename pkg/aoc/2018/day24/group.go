package day24

import "strings"

type group struct {
	n, hp     int    // unit count, hitpoints
	str, init int    // attack strength, initiative
	id        int    // group id
	dtype     string // damage type
	weak      string // weaknesses
	immune    string // immunities
}

// ep returns the effective power of a group
func (g *group) ep() int {
	return g.n * g.str
}

func (g *group) damage(target *group) int {
	if strings.Contains(target.immune, g.dtype) {
		return 0
	}
	if strings.Contains(target.weak, g.dtype) {
		return g.ep() * 2
	}
	return g.ep()
}
