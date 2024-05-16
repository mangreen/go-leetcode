package q683_KEmptySlots

import (
	"math"
)

func kEmptySlots(bulbs []int, k int) int {
	days := make([]int, len(bulbs))
	for i, b := range bulbs {
		days[b-1] = i+1
	}

	res := math.MaxInt32

	l := 0
	r := k+1

	for i:=0; r<len(days); i++ {
		if i==r {
			res = min(res, max(days[l], days[r]))
			l = i
			r = l+k+1

			continue
		}

		if days[i]<days[l] || days[i]<days[r] {
			l = i
			r = l+k+1

			continue
		}
	}

	if res == math.MaxInt32 {
		return -1
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
/*
bulbs=[1, 3, 2]
k=1

	  i->day0
days=[1, 3, 2]
	  |
	  position1

	  i
days=[1, 3, 2]
l=0 -> i=0
r=k+1=2 != i=0 -> l+k+1=0+1+1=2

	     i
days=[1, 3, 2]
l=0
r=2 != i=1

	        i
days=[1, 3, 2]
l=0
r=2 == i=2
res=min(res, max(2, 2))=2
*/