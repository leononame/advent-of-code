package points

import "gitlab.com/leononame/advent-of-code-2018/pkg/geo"

func Neighbours(p geo.Pointer) []geo.Pointer {
	return []geo.Pointer{p.Down(), p.Left(), p.Up(), p.Right()}
}
