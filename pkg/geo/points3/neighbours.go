package points3

import "gitlab.com/leononame/advent-of-code-2018/pkg/geo"

func Neighbours(p geo.Pointer3) []geo.Pointer3 {
	return []geo.Pointer3{p.Down(), p.Left(), p.Up(), p.Right(), p.Higher(), p.Lower()}
}
