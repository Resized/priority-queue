package priority_queue

type entry[T any] struct {
	key int
	val T
}

func newEntry[T any](key int, val T) entry[T] {
	return entry[T]{key: key, val: val}
}

type PriorityQueue[T any] interface {
	Push(key int, val T)
	Pop() T
	Peek() T
	IsEmpty() bool
	Len() int
}

type priorityQueue[T any] struct {
	isMinPQ bool
	heap    []entry[T]
}

func NewMinPQ[T any]() PriorityQueue[T] {
	return &priorityQueue[T]{isMinPQ: true}
}

func NewMaxPQ[T any]() PriorityQueue[T] {
	return &priorityQueue[T]{isMinPQ: false}
}

func NewMinPQFromArr[T any](a []entry[T]) PriorityQueue[T] {
	pq := &priorityQueue[T]{heap: a, isMinPQ: true}
	pq.init()
	return pq
}

func NewMaxPQFromArr[T any](a []entry[T]) PriorityQueue[T] {
	pq := &priorityQueue[T]{heap: a, isMinPQ: false}
	pq.init()
	return pq
}

func (pq *priorityQueue[T]) init() {
	// heapify
	n := len(pq.heap)
	for i := n/2 - 1; i >= 0; i-- {
		pq.down(i, n)
	}
}

func (pq *priorityQueue[T]) IsEmpty() bool {
	return len(pq.heap) == 0
}

func (pq *priorityQueue[T]) Push(key int, value T) {
	e := newEntry(key, value)
	pq.heap = append(pq.heap, e)
	pq.up(len(pq.heap) - 1)
}

func (pq *priorityQueue[T]) Len() int {
	return len(pq.heap)
}

func (pq *priorityQueue[T]) Pop() T {
	n := len(pq.heap) - 1
	pq.heap[0], pq.heap[n] = pq.heap[n], pq.heap[0]
	pq.down(0, n)
	x := pq.heap[n]
	pq.heap[n] = entry[T]{}
	pq.heap = pq.heap[:n]
	return x.val
}

func (pq *priorityQueue[T]) Peek() T {
	return pq.heap[0].val
}

func (pq *priorityQueue[T]) up(current int) {
	for {
		parent := (current - 1) / 2 // parent
		cur, next := pq.heap[current].key, pq.heap[parent].key
		if current == parent || ((pq.isMinPQ && cur >= next) || (!pq.isMinPQ && cur <= next)) {
			break
		}
		pq.heap[parent], pq.heap[current] = pq.heap[current], pq.heap[parent]
		current = parent
	}
}

func (pq *priorityQueue[T]) down(startIndex int, n int) bool {
	current := startIndex
	for {
		leftChild := 2*current + 1
		rightChild := leftChild + 1
		if leftChild >= n || leftChild < 0 { // leftChild < 0 after int overflow
			break
		}
		child := leftChild // left child
		r, l := pq.heap[rightChild].key, pq.heap[leftChild].key
		if rightChild < n && ((pq.isMinPQ && r < l) || (!pq.isMinPQ && r > l)) {
			child = rightChild
		}
		cur, next := pq.heap[current].key, pq.heap[child].key
		if (pq.isMinPQ && cur <= next) || (!pq.isMinPQ && cur >= next) {
			break
		}
		pq.heap[current], pq.heap[child] = pq.heap[child], pq.heap[current]
		current = child
	}
	return current > startIndex
}
