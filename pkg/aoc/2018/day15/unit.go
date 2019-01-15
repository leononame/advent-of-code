package day15

import (
	"github.com/pkg/errors"

	"gitlab.com/leononame/advent-of-code-2018/pkg/geo"
)

type unit struct {
	hp      int
	pow     int
	pos     geo.Point
	t       rune
	enemies *[]*unit
	cave    *cave
	id      int
}

type target struct {
	pos      geo.Point
	enemy    *unit
	distance int
}

func (u *unit) tick() bool {
	if e := u.enemyInReach(); e != nil {
		return u.attack(e)
	}
	targets := u.findTargets()
	t, err := u.selectTarget(targets)
	// no target could be pathed to, exit
	if err != nil {
		return false
	}
	u.approach(t)
	if e := u.enemyInReach(); e != nil {
		return u.attack(e)
	}
	return false
}

func (u *unit) attack(enemy *unit) bool {
	enemy.hp -= u.pow
	if enemy.hp <= 0 {
		enemy.hp = 0
		return true
	}
	return false
}

func (u *unit) enemy(p geo.Point) *unit {
	for _, e := range *u.enemies {
		if e.pos == p {
			return e
		}
	}
	return nil
}

func (u *unit) selectEnemy(e1, e2 *unit) *unit {
	if e1 == nil {
		return e2
	} else if e2 == nil {
		return e1
	}
	if e1.hp < e2.hp || (e1.hp == e2.hp && readingOrder(e1.pos, e2.pos)) {
		return e1
	}
	return e2
}

func (u *unit) enemyInReach() *unit {
	var selected *unit
	for _, n := range u.pos.Neighbours() {
		selected = u.selectEnemy(selected, u.enemy(n))
	}
	return selected
}

func (u *unit) findTargets() []target {
	var ts []target
	for _, e := range *u.enemies {
		for _, n := range u.cave.neighbours(e.pos) {
			ts = append(ts, target{n, e, 0})
		}
	}
	return ts
}

func (u *unit) selectTarget(ts []target) (target, error) {
	best := target{distance: 10000, pos: geo.Point{}}
	for _, t := range ts {
		t.path(u.pos, best.distance)
		if t.distance < best.distance ||
			(t.distance == best.distance && readingOrder(t.pos, best.pos)) {
			best = t
		}
	}
	if best.distance == 10000 {
		return best, errors.New("Target not found")
	}
	return best, nil
}

func (u *unit) approach(t target) *unit {
	best := t.distance - 1
	loc := u.pos.Down()
	for _, n := range u.cave.neighbours(u.pos) {
		t.path(n, best)
		if t.distance == best && readingOrder(n, loc) {
			loc = n
		}
	}
	return u.moveTo(loc)

	// if t.path(u.pos.Up(), best); t.distance == best {
	// 	return u.moveTo(u.pos.Up())
	// } else if t.path(u.pos.Left(), best); t.distance == best {
	// 	return u.moveTo(u.pos.Left())
	// } else if t.path(u.pos.Right(), best); t.distance == best {
	// 	return u.moveTo(u.pos.Right())
	// }
	// return u.moveTo(u.pos.Left())
}

func (u *unit) moveTo(p geo.Point) *unit {
	u.cave.layout[u.pos.GetY()][u.pos.GetX()] = floor
	u.cave.layout[p.GetY()][p.GetX()] = u.t
	u.pos = p
	return u
}
