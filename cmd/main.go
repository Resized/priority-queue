package main

import "github.com/Resized/priority-queue/pkg/priority-queue"

func main() {
	pq := priority_queue.NewMinPQ[string]()
	pq.Push(1, "This ")
	pq.Push(7, "sentence")
	pq.Push(4, "complete ")
	pq.Push(2, "is ")
	pq.Push(3, "a ")
	pq.Push(9, "!")

	var result string
	for !pq.IsEmpty() {
		result += pq.Pop()
	}
	println(result)

	pq2 := priority_queue.NewMinPQ[int]()
	pq2.Push(1, 1)
	pq2.Push(1, 2)
	pq2.Push(1, 3)
	pq2.Push(1, 4)
	pq2.Push(1, 5)
	pq2.Push(1, 6)

	for !pq2.IsEmpty() {
		println(pq2.Pop())
	}

	m := map[uint64]string{1: "This ", 7: "sentence", 5: "complete ", 2: "is ", 3: "also ", 9: "!", 4: "a "}
	pq3 := priority_queue.NewMinPQFromMap(m)
	var result2 string
	for !pq3.IsEmpty() {
		result2 += pq3.Pop()
	}
	println(result2)
}
