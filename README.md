Blocking Queues
---
Blocking Queues provides some simple, performant, goroutine safe queues useful as resource pools or job queues. 
The primary focus is simplicity and high performance without sacrificing readability. In fact I tried to
provide good documentation on the code and some examples of usage.


## Queues Provided
* **ArrayBlockingQueue**: A bounded blocking queue backed by a slice
* **ArrayBlockingDeque**: A bounded blocking deque backed by a slice
* **ArrayBlockingPriorityQueue**: A bounded blocking PriorityQueue backed by a slice

* **LinkedBlockingQueue**: A bounded blocking queue backed by a container/list
* **LinkedBlockingDeque**: A bounded blocking deque backed by a container/list

## Installation
```go
go get -u github.com/theodesp/blockingQueues
```

## Usage
TODO

## Benchmarks
TODO

## LICENCE
Copyright © 2017 Theo Despoudis MIT license
