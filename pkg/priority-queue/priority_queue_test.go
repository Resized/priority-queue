package priority_queue

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestMinPriorityQueue_Push(t *testing.T) {
	pq := NewMinPQ[string]()
	table := []struct {
		key   int
		value string
	}{
		{1, "This "},
		{7, "sentence"},
		{4, "complete "},
		{2, "is "},
		{3, "a "},
		{9, "!"},
	}
	for _, tt := range table {
		pq.Push(tt.key, tt.value)
	}
	assert.Equal(t, 6, pq.Len())

	var result string
	for range table {
		result += pq.Pop()
	}
	assert.Equal(t, "This is a complete sentence!", result)
}

func TestMaxPriorityQueue_Push(t *testing.T) {
	pq := NewMaxPQ[string]()
	table := []struct {
		key   int
		value string
	}{
		{9, "This "},
		{2, "sentence"},
		{4, "complete "},
		{7, "is "},
		{5, "a "},
		{1, "!"},
	}
	for _, tt := range table {
		pq.Push(tt.key, tt.value)
	}
	assert.Equal(t, 6, pq.Len())

	var result string
	for range table {
		result += pq.Pop()
	}
	assert.Equal(t, "This is a complete sentence!", result)
}

func TestMaxPriorityQueue_IsEmpty(t *testing.T) {
	pq := NewMaxPQ[string]()
	assert.True(t, pq.IsEmpty())
	pq.Push(1, "Hello")
	assert.False(t, pq.IsEmpty())
	pq.Pop()
	assert.True(t, pq.IsEmpty())
}

func TestMinPriorityQueue_IsEmpty(t *testing.T) {
	pq := NewMinPQ[string]()
	assert.True(t, pq.IsEmpty())
	pq.Push(1, "Hello")
	assert.False(t, pq.IsEmpty())
	pq.Pop()
	assert.True(t, pq.IsEmpty())
}

func TestNewMinPQFromMap(t *testing.T) {
	m := map[int]string{1: "This ", 7: "sentence", 4: "complete ", 2: "is ", 3: "a ", 9: "!"}
	pq := NewMinPQFromMap(m)
	assert.Equal(t, 6, pq.Len())

	var result string
	for range m {
		result += pq.Pop()
	}
	assert.Equal(t, "This is a complete sentence!", result)
}

func TestNewMaxPQFromMap(t *testing.T) {
	m := map[int]string{9: "This ", 2: "sentence", 4: "complete ", 7: "is ", 5: "a ", 1: "!"}
	pq := NewMaxPQFromMap(m)
	assert.Equal(t, 6, pq.Len())

	var result string
	for range m {
		result += pq.Pop()
	}
	assert.Equal(t, "This is a complete sentence!", result)
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
func initpq(b *testing.B) PriorityQueue[int] {
	pq := NewMinPQ[int]()
	for i := 0; i < b.N; i++ {
		pq.Push(rand.Int(), 5)
	}
	return pq
}
