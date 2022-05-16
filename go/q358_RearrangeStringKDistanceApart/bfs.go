package q159_LongestSubstringWithAtMost2DistinctCharacters

// import "container/heap"

// func rearrangeString(s string, k int) string {
// 	if k == 0 {
// 		return s
// 	}

// 	ans := ""

// 	_l := len(s)

// 	m := make(map[byte]int)
// 	for i := range s {
// 		m[s[i]]++
// 	}

// 	h := &IntHeap{}
// 	heap.Init(h)

// 	for k, v := range m {
// 		heap.Push(h, k-v)
// 	}
// }

// 建立自訂型態
type IntHeap []int

// 實作以滿足 sort.Interface 的部分
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	// Push 和 Pop 要透過指標傳值，因為會更動到 slice 的長度而不只是內容
	// 解釋：
	// append 過後的切片可能已經不是本來的切片
	// 所以要透過指標更改 h 的值
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	// 先將要回傳的值暫存起來
	x := old[n-1]
	// 切片重切，因為 heap 套件在內部實作時是靠
	// 切片的長度在定位的，跟我們先前的方式不太一樣
	// 所以一定要重切
	*h = old[0 : n-1]
	return x
}
