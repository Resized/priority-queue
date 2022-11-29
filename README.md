# Priority Queue 
A simple to use priority queue using a binary heap implemented in GO.

- Using generics which means it supports any declared data type with safety
- Safe for concurrency.
- Inserting multiple items with the same priority will  
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
The time complexity of all operations is O(log N)
### Time Complexity

- Write -  O(1)
- Read - O(K)

### Space Complexity
- O(N+K)

## Optimization
The time complexity of the reading is always O(K) since we are required to go through all the lists 
in order from the first priority to the last, however I found a simple solution to optimize it by 
adding the firstAvailable parameter which saves the first available list.

This solution reduces read time to O(1) in best case scenario, and statistically it should lower the 
overall read time on average.
