package priority_queue

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPriorityQueue_Push(t *testing.T) {
	pq := New[string]()
	pq.Push(7, "b")
	pq.Push(2, "a")
	pq.Push(9, "c")
	assert.EqualValues(t, "a", pq.Pop())
	assert.EqualValues(t, "b", pq.Pop())
	assert.EqualValues(t, "c", pq.Pop())
}

func BenchmarkPriorityQueue_Push(b *testing.B) {
	b.StopTimer()
	pq := initpq(b)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(rand.Int(), 5)
	}
}

func BenchmarkPriorityQueue_Pop(b *testing.B) {
	b.StopTimer()
	pq := initpq(b)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}
func initpq(b *testing.B) *PriorityQueue[int] {
	pq := New[int]()
	for i := 0; i < b.N; i++ {
		pq.Push(rand.Int(), 5)
	}
	return pq
}
