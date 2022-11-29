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

	print(pq.Pop())
	print(pq.Pop())
	print(pq.Pop())
	print(pq.Pop())
	print(pq.Pop())
	print(pq.Pop())

}
