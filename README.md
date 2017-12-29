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

```text
ArrayBlockingQueueSuite.BenchmarkPeek     100000000               21.0 ns/op
ArrayBlockingQueueSuite.BenchmarkPop      100000000               20.7 ns/op
ArrayBlockingQueueSuite.BenchmarkPopOverflow       100000000               20.8 ns/op
ArrayBlockingQueueSuite.BenchmarkPush      50000000                38.9 ns/op
ArrayBlockingQueueSuite.BenchmarkPushOverflow      50000000                39.0 ns/op
```

Multiple Goroutines - Different ratio of readers/writers

```text
PASS: arrayBlockingQueue_test.go:221: ArrayBlockingQueueSuite.BenchmarkPut1to1  10000000               177 ns/op
```

## LICENCE
Copyright © 2017 Theo Despoudis MIT license
