package main

import (
	"priority-queue/pkg/priority-queue"
)

func main() {
	pq := priority_queue.New[string]()
	pq.Push(7, "World")
	pq.Push(2, "Hello ")
	pq.Push(9, "!")
	print(pq.Pop())
	print(pq.Pop())
	print(pq.Pop())

}
