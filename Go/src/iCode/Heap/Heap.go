package Heap

/**
@Description: 堆(大顶堆)
@Date: 1/30/2021
@Author: lichang
*/

type Heap struct {
	a     []int // 堆顶元素在a[1]
	cap   int
	count int
}

// init heap
func NewHeap(capacity int) *Heap {
	return &Heap{
		a:     make([]int, capacity),
		cap:   capacity,
		count: 0,
	}
}

// 大顶堆调整
func (h *Heap) Insert(data int) {
	if h.count == h.cap { // 堆满
		return
	}

	h.count++
	h.a[h.count] = data

	// 调整堆
	i := h.count
	parent := i / 2
	for parent > 0 && h.a[parent] < h.a[i] { // 未比较到堆顶且父节点小于子节点
		swap(h.a, parent, i)
		i = parent
		parent = i / 2
	}
}

func (h *Heap) removeMax() {
	if h.count == 0 {
		return
	}
	swap(h.a, 1, h.count)
	h.count--

	HeapifyUp2Down(h.a, h.count)
}

func HeapifyUp2Down(a []int, count int) {
	for i := 1; i <= count/2; {
		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}
		if maxIndex == i {
			break
		}
		swap(a, i, maxIndex)
		i = maxIndex
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
