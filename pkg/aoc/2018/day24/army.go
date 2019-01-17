package day24

import (
	"sort"
)

type army []*group

func (a *army) chooseTarget(enemies army) []*fight {
	// Sort by attack strength, then initiative
	sort.Sort(byStrInit(*a))
	var fs []*fight
	remaining := make(army, len(enemies))
	copy(remaining, enemies)
	for _, g := range *a {
		max, enemy, idx := 0, &group{}, 0
		for i, e := range remaining {
			d := g.damage(e)

			if d > max ||
				(d == max && (e.ep() > enemy.ep() ||
					(e.ep() == enemy.ep() && e.init > enemy.init))) {
				max = d
				enemy = e
				idx = i
			} else if d == max {

			}
		}
		logger.Debugf("fighter ID: %2d, idx: %2d, enemy ID: %2d, damage: %9d, data: ", g.id, idx, enemy.id, max)
		if max == 0 {
			continue
		}
		fs = append(fs, &fight{g, enemy})
		remaining = append(remaining[0:idx], remaining[idx+1:len(remaining)]...)
	}
	return fs
}

func (a *army) removeDead() {
	i := 0
	for _, g := range *a {
		if g.n == 0 {
			continue
		}
		(*a)[i] = g
		i++
	}
	*a = (*a)[:i]
}

func (a *army) countUnits() int {
	sum := 0
	for _, g := range *a {
		sum += g.n
	}
	return sum
}

// byStrInit is a type to sort an army by strength and then initiative
type byStrInit army

func (b byStrInit) Len() int {
	return len(b)
}

func (b byStrInit) Less(i int, j int) bool {
	// effective power of groups i and j
	epi, epj := b[i].ep(), b[j].ep()
	return epi > epj || (epi == epj && b[i].init > b[j].init)
}

func (b byStrInit) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}
