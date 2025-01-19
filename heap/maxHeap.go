package heap

// MaxHeapPosition contains the position in x, y coordinate and the position value
type MaxHeapPosition struct {
	x     int
	y     int
	value int
}

type MaxHeap []MaxHeapPosition

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value >= h[j].value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(MaxHeapPosition))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}