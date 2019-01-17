package day24

import (
	"sort"
)

type fight struct {
	attacker, defender *group
}

func (f *fight) calc() {
	if f.attacker.n == 0 {
		return
	}
	// Need to recalc damage because units might have been lost
	damage := f.attacker.damage(f.defender)
	unitsKilled := damage / f.defender.hp
	f.defender.n -= unitsKilled
	if f.defender.n < 0 {
		f.defender.n = 0
	}
}

type round []*fight

func (r round) Len() int {
	return len(r)
}

func (r round) Less(i int, j int) bool {
	return r[i].attacker.init > r[j].attacker.init
}

func (r round) Swap(i int, j int) {
	r[i], r[j] = r[j], r[i]
}

type battle struct {
	imm, inf army
	all      army
}

func (b *battle) round() {
	// Get fight combinations
	fights := b.inf.chooseTarget(b.imm)
	fights = append(fights, b.imm.chooseTarget(b.inf)...)
	r := round(fights)
	// Sort all units by initiative
	sort.Sort(r)
	// Now do the fighting
	for _, f := range r {
		f.calc()
	}
	b.removeDead()
}

func (b *battle) removeDead() {
	b.imm.removeDead()
	b.inf.removeDead()
	b.all.removeDead()
}

func (b *battle) fight() *army {
	remainingUnits := b.imm.countUnits() + b.inf.countUnits()
	for {
		b.round()
		if len(b.inf) == 0 {
			return &b.imm
		} else if len(b.imm) == 0 {
			return &b.inf
		}
		tmp := b.imm.countUnits() + b.inf.countUnits()
		if remainingUnits == tmp {
			logger.Debug("Deadlock")
			return nil
		}
		remainingUnits = tmp
	}

}
