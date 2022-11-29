# Priority Queue 
A simple to use priority queue using a binary heap implemented in GO.

- Using generics which means it supports any declared data type with safety.
- Safe for concurrency.
- Retrieving items with the same priority will retrieve them in FIFO order.
## Example

### Import

```go    
import "github.com/Resized/priority-queue/pkg/priority-queue"
```

### Usage

```go
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

```
#### Outputs
```
This is a complete sentence!
```

## Complexity

### Time Complexity

- Push -  O(log N)
- Pop - O(log N)
- Peek - O(1)

### Space Complexity
- O(N)
