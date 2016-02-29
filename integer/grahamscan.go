package integer

import "sort"

type convexHullSortByPolarAngle SliceTuple

func (spa convexHullSortByPolarAngle) Less(i, j int) bool {
	i++
	j++
	ccw := SliceTuple(spa).PointsCCW(0, i, j)
	if ccw != 0 {
		return ccw > 0
	}

	// sort collinears from furthest away to nearest
	dx := Abs((*spa[j])[0]-(*spa[0])[0]) - Abs((*spa[i])[0]-(*spa[0])[0])
	if dx != 0 {
		return dx < 0
	}

	return ((*spa[j])[1]-(*spa[0])[1])-((*spa[i])[1]-(*spa[0])[1]) < 0
}
func (spa convexHullSortByPolarAngle) Swap(i, j int) {
	i++
	j++
	spa[i], spa[j] = spa[j], spa[i]
}
func (spa convexHullSortByPolarAngle) Len() int { return len(spa) - 1 }

func (st SliceTuple) PointsCCW(i, j, k int) int {
	// (j.x - i.x)*(k.y - i.y) - (j.y - i.y)*(k.x - i.x)
	return (((*st[j])[0]-(*st[i])[0])*((*st[k])[1]-(*st[i])[1]) - ((*st[j])[1]-(*st[i])[1])*((*st[k])[0]-(*st[i])[0]))
}

func (st SliceTuple) PointsDistance2(i, j int) int {
	dx := ((*st[j])[0] - (*st[i])[0])
	dy := ((*st[j])[1] - (*st[i])[1])
	return dx*dx + dy*dy
}

func (st SliceTuple) convexHullLowestY() int {
	mini := 0
	miny := (*st[0])[1]
	minx := (*st[0])[0]
	for i := 1; i < len(st); i++ {
		y := (*st[i])[1]
		x := (*st[i])[0]
		if y < miny {
			mini = i
			miny = y
			minx = x
			continue
		}
		if y == miny && x < minx {
			mini = i
			miny = y
			minx = x
		}
	}
	return mini
}

func (st SliceTuple) ConvexHull(includeCollinears bool) int {
	if len(st) <= 2 {
		return len(st)
	}

	st.Swap(0, st.convexHullLowestY())

	sort.Sort(convexHullSortByPolarAngle(st))

	M := 1
	i := 2
	end := len(st)

	// skip or add beginning collinears
	for st.PointsCCW(0, 1, i) == 0 {
		i++
		if includeCollinears {
			M++
		}
		if i == end {
			break
		}
	}
	if includeCollinears {
		st[1 : M+1].Reverse()
	}
	if i == end {
		return M + 1
	}

	// skip ending collinears
	for st.PointsCCW(0, end-1, end-2) == 0 {
		end--
	}

	for ; i < end; i++ {
		if includeCollinears {
			for st.PointsCCW(M-1, M, i) < 0 {
				M--
			}
		} else {
			for st.PointsCCW(M-1, M, i) <= 0 {
				M--
			}
		}

		M++
		st.Swap(M, i)
	}

	// add ending collinears
	if includeCollinears {
		for i := end; i < len(st); i++ {
			M++
			st.Swap(M, i)
		}
	}
	return M + 1
}
