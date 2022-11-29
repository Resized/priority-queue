package priority_queue

type entry[T any] struct {
	key int
	val T
}

func newEntry[T any](key int, val T) entry[T] {
	return entry[T]{key: key, val: val}
}

type PriorityQueue[T any] struct {
	heap []entry[T]
}

func New[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

func NewFromArr[T any](a []entry[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{heap: a}
	pq.init()
	return pq
}

func (pq *PriorityQueue[T]) init() {
	// heapify
	n := len(pq.heap)
	for i := n/2 - 1; i >= 0; i-- {
		pq.down(i)
	}
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.heap) == 0
}

func (pq *PriorityQueue[T]) Push(key int, value T) {
	e := newEntry(key, value)
	pq.heap = append(pq.heap, e)
	pq.up(len(pq.heap) - 1)
}

func (pq *PriorityQueue[T]) Pop() T {
	n := len(pq.heap) - 1
	pq.heap[0], pq.heap[n] = pq.heap[n], pq.heap[0]
	pq.down(0)
	x := pq.heap[n]
	pq.heap[n] = entry[T]{}
	pq.heap = pq.heap[:n]
	return x.val
}

func (pq *PriorityQueue[T]) Peek() T {
	return pq.heap[0].val
}

func (pq *PriorityQueue[T]) up(current int) {
	for {
		next := (current - 1) / 2 // parent
		if pq.heap[next].key <= pq.heap[current].key {
			break
		}
		pq.heap[next], pq.heap[current] = pq.heap[current], pq.heap[next]
		current = next
	}
}

func (pq *PriorityQueue[T]) down(startIndex int) bool {
	current := startIndex
	n := len(pq.heap)
	for {
		leftChild := 2*current + 1
		if leftChild >= n || leftChild < 0 { // leftChild < 0 after int overflow
			break
		}
		next := leftChild           // left child
		rightChild := leftChild + 1 // right child
		if rightChild < n && pq.heap[rightChild].key < pq.heap[leftChild].key {
			next = rightChild
		}
		if pq.heap[current].key <= pq.heap[next].key {
			break
		}
		pq.heap[current], pq.heap[next] = pq.heap[next], pq.heap[current]
		current = next
	}
	return current > startIndex
}
