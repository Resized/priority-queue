package priority_queue

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestMinPriorityQueue_Push(t *testing.T) {
	pq := NewMinPQ[string]()
	table := []struct {
		key   uint64
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

func TestMinPriorityQueue_PushParallel(t *testing.T) {
	pq := NewMinPQ[string]()
	numIterations := 10000
	done := make(chan struct{})
	table := []struct {
		key   uint64
		value string
	}{
		{1, "This "},
		{7, "sentence"},
		{4, "complete "},
		{2, "is "},
		{3, "a "},
		{9, "!"},
	}
	for i := 0; i < numIterations; i++ {
		go func() {
			for _, tt := range table {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
				pq.Push(tt.key, tt.value)
			}
			done <- struct{}{}
		}()
	}
	for i := 0; i < numIterations; i++ {
		<-done
	}
	assert.Equal(t, 6*numIterations, pq.Len())
	var expected string
	var result string
	for range table {
		for i := 0; i < numIterations; i++ {
			result += pq.Pop()
		}
	}
	expectedOrder := []string{"This ", "is ", "a ", "complete ", "sentence", "!"}
	for _, tt := range expectedOrder {
		for i := 0; i < numIterations; i++ {
			expected += tt
		}
	}
	assert.Equal(t, expected, result)
}

func TestMinPriorityQueueTableDriven(t *testing.T) {
	type e struct {
		key uint64
		val string
	}
	type res struct {
		str string
		len int
	}
	table := []struct {
		name     string
		data     []e
		expected res
	}{
		{name: "random", data: []e{{1, "This "}, {7, "sentence"}, {4, "complete "}, {2, "is "}, {3, "a "}, {9, "!"}}, expected: res{"This is a complete sentence!", 6}},
		{name: "same priority", data: []e{{1, "This "}, {1, "is "}, {1, "a "}, {1, "complete "}, {1, "sentence"}, {1, "!"}}, expected: res{"This is a complete sentence!", 6}},
		{name: "empty", data: []e{}, expected: res{"", 0}},
		{name: "one element", data: []e{{1, "This "}}, expected: res{"This ", 1}},
		{name: "two elements", data: []e{{7, "is "}, {4, "This "}}, expected: res{"This is ", 2}},
		{name: "reverse order", data: []e{{9, "!"}, {7, "sentence"}, {4, "complete "}, {3, "a "}, {2, "is "}, {1, "This "}}, expected: res{"This is a complete sentence!", 6}},
	}
	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewMinPQ[string]()
			for _, d := range tt.data {
				pq.Push(d.key, d.val)
			}
			assert.Equal(t, tt.expected.len, pq.Len())

			var result string
			for range tt.data {
				result += pq.Pop()
			}
			assert.Equal(t, tt.expected.str, result)
		})
	}
}

func TestMinPriorityQueue_PushSamePriority(t *testing.T) {
	pq := NewMinPQ[string]()
	table := []struct {
		key   uint64
		value string
	}{
		{1, "This "},
		{1, "is "},
		{1, "a "},
		{1, "complete "},
		{1, "sentence"},
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

func TestMaxPriorityQueue_Push(t *testing.T) {

	pq := NewMaxPQ[string]()
	table := []struct {
		key   uint64
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
	m := map[uint64]string{1: "This ", 7: "sentence", 4: "complete ", 2: "is ", 3: "a ", 9: "!"}
	pq := NewMinPQFromMap(m)
	assert.Equal(t, 6, pq.Len())

	var result string
	for range m {
		result += pq.Pop()
	}
	assert.Equal(t, "This is a complete sentence!", result)
}

func TestNewMaxPQFromMap(t *testing.T) {
	m := map[uint64]string{9: "This ", 2: "sentence", 4: "complete ", 7: "is ", 5: "a ", 1: "!"}
	pq := NewMaxPQFromMap(m)
	assert.Equal(t, 6, pq.Len())

	var result string
	for range m {
		result += pq.Pop()
	}
	assert.Equal(t, "This is a complete sentence!", result)
}

func BenchmarkPriorityQueue_Push(b *testing.B) {
	q := NewMinPQ[int]()
	for i := 0; i < b.N; i++ {
		q.Push(uint64(rand.Int()), rand.Int())
	}
}

func BenchmarkPriorityQueue_Pop(b *testing.B) {
	q := NewMinPQ[int]()
	for i := 0; i < b.N; i++ {
		q.Push(uint64(rand.Int()), rand.Int())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}
