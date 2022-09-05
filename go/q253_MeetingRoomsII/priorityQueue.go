package q252_MeetingRooms

import (
	"container/heap"
	"sort"
)

func minMeetingRooms2(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &IntervalMinHeap{}
	heap.Init(h)

	for _, v := range intervals {
		if h.Len() == 0 {
			heap.Push(h, Interval{
				Start: v[0],
				End: v[1],
			})

			continue
		}

		top := (*h)[0]

		if top.End < v[0] {
			heap.Pop(h)
		}

		heap.Push(h, Interval{
			Start: v[0],
			End: v[1],
		})
	}

	return h.Len()
}

type Interval struct {
	Start int
	End   int
}

type IntervalMinHeap []Interval

func (h IntervalMinHeap) Len() int {
	return len(h)
}

func (h IntervalMinHeap) Less(i, j int) bool {
	return h[i].End < h[j].End
}

func (h IntervalMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntervalMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Interval))
}

func (h *IntervalMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
