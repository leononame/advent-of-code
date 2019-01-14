package day22

type Tile struct {
	Region
	tool int
}

type TraverseMap struct {
	*Map
}

func (m *TraverseMap) searchNeighbours(i Tile) []Tile {
	var ns []Tile
	for _, r := range m.neighbours(i.Location) {
		// Allowed tools as bitmap
		allowed := tools(r.Terrain)
		// If our current tool is allowed to enter next tile, create new neighbour
		if allowed&i.tool == i.tool {
			ns = append(ns, Tile{Region: *r, tool: i.tool})
		}
		// Get the other tool by toggling our current tool
		other := tools(i.Terrain) ^ i.tool
		// If the other tool is allowed to enter, create new neighbour
		// (The coordinate is the same, but the tool is different)
		if allowed&other == other {
			ns = append(ns, Tile{Region: i.Region, tool: other})
		}
	}
	return ns
}
