package main

import (
	"priority-queue/pkg/priority-queue"
)

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

	m := map[int]string{1: "This ", 7: "sentence", 5: "complete ", 2: "is ", 3: "also ", 9: "!", 4: "a "}
	pq2 := priority_queue.NewMinPQFromMap(m)
	var result2 string
	for !pq2.IsEmpty() {
		result2 += pq2.Pop()
	}
	println(result2)
}
