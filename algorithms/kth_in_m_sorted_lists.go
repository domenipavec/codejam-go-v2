package algorithms

import (
	"log"
	"sort"
)

type listLen func(list int) int
type listLess func(list1, index1, list2, index2 int) bool

// TODO
// this is implmented by https://stackoverflow.com/a/8754265 but does not work
// i think algorithm is faulty
// this would be used for 2018/kpractice/d
func KthInMSortedLists(M int, ListLen listLen, Less listLess, k int) (int, int) {
	listStarts := make([]int, M)
	listEnds := make([]int, M)
	for list := range listEnds {
		listEnds[list] = ListLen(list)
	}

	type listMiddle struct {
		List  int
		Index int
	}
	middles := make([]listMiddle, M)
	listMiddleLess := func(i, j int) bool {
		return Less(
			middles[i].List, listStarts[middles[i].List]+middles[i].Index,
			middles[j].List, listStarts[middles[j].List]+middles[j].Index,
		)
	}

	for {
		maxLen := 0
		middles = middles[:0]
		for list := range listStarts {
			length := listEnds[list] - listStarts[list]
			if length > maxLen {
				maxLen = length
			}

			if length < 1 {
				continue
			}

			middles = append(middles, listMiddle{
				List:  list,
				Index: (length - 1) / 2,
			})
		}
		sort.Slice(middles, listMiddleLess)
		log.Println(listStarts, listEnds, k)
		log.Println(middles)
		if maxLen <= 1 {
			break
		}

		indexSum := 0
		for _, middle := range middles {
			indexSum += middle.Index + 1
			if indexSum <= k {
				listStarts[middle.List] += middle.Index + 1
				k -= middle.Index + 1
				indexSum -= middle.Index + 1
			} else {
				listEnds[middle.List] = listStarts[middle.List] + middle.Index + 1
			}
		}
	}

	indexSum := 0
	for _, middle := range middles {
		indexSum += middle.Index + 1
		if indexSum > k {
			return middle.List, listStarts[middle.List] + middle.Index
		}
	}

	panic("something went wrong")
}
