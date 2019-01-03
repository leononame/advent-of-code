package points3

import "gitlab.com/leononame/advent-of-code-2018/pkg/geo3"

func Neighbours(p geo3.Pointer) []geo3.Pointer {
	return []geo3.Pointer{p.Down(), p.Left(), p.Up(), p.Right(), p.Higher(), p.Lower()}
}
