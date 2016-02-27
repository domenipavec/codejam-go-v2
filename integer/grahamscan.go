package integer

import "sort"

type convexHullSortByPolarAngle SliceTuple

func (spa convexHullSortByPolarAngle) Less(i, j int) bool {
	i++
	j++
	return SliceTuple(spa).PointsCCW(0, i, j) > 0
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

func (st SliceTuple) ConvexHull() int {
	if len(st) <= 3 {
		return len(st)
	}

	st.Swap(0, st.convexHullLowestY())

	sort.Sort(convexHullSortByPolarAngle(st))

	M := 1
	for i := 2; i < len(st); i++ {
		for st.PointsCCW(M-1, M, i) < 0 {
			M--
		}

		M++
		st.Swap(M, i)

		if st.PointsCCW(M-2, M-1, M) == 0 {
			if st.PointsDistance2(M-2, M) > st.PointsDistance2(M-2, M-1) {
				st.Swap(M, M-1)
			}
			M--
		}
	}
	return M + 1
}
