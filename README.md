Blocking Queues
---
Blocking Queues provides some simple, performant, goroutine safe queues useful as resource pools or job queues. 
The primary focus is simplicity and high performance without sacrificing readability. In fact I tried to
provide good documentation on the code and some examples of usage.


## Queues Provided
* **ArrayBlockingQueue**: A bounded blocking queue backed by a slice
* **LinkedBlockingQueue**: A bounded blocking queue backed by a container/list

## Installation
```go
go get -u github.com/theodesp/blockingQueues
```

## Usage
TODO

## Benchmarks
Using:
  ```text
  Model Name:	MacBook Pro
  Model Identifier:	MacBookPro12,1
  Processor Name:	Intel Core i7
  Processor Speed:	3.1 GHz
  Number of Processors:	1
  Total Number of Cores:	2
  L2 Cache (per Core):	256 KB
  L3 Cache:	4 MB
  Memory:	16 GB
```

Simple operations - no goroutines

```go
PASS: ArrayBlockingQueueSuite.BenchmarkPop    50000000                49.1 ns/op
PASS: ArrayBlockingQueueSuite.BenchmarkPopOverflow    50000000                51.3 ns/op
PASS: ArrayBlockingQueueSuite.BenchmarkPush   20000000                84.7 ns/op
PASS: ArrayBlockingQueueSuite.BenchmarkPushOverflow   20000000                75.3 ns/op
```

## LICENCE
Copyright © 2017 Theo Despoudis MIT license
