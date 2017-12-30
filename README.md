Blocking Queues
---
<a href="https://godoc.org/github.com/theodesp/blockingQueues">
<img src="https://godoc.org/github.com/theodesp/blockingQueues?status.svg" alt="GoDoc">
</a>

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

Multiple Goroutines - Different ratio of readers/writers - Size of Queue is 1024 items

```text
ArrayBlockingQueueSuite.BenchmarkPut1to1   10000000               169 ns/op
ArrayBlockingQueueSuite.BenchmarkPut2to2    5000000               508 ns/op
ArrayBlockingQueueSuite.BenchmarkPut4to4    1000000              1222 ns/op
```
The last test is slower as the number of goroutines are double of the available logic cores
and it is expected to go slower because of the context switching.

```text
ArrayBlockingQueueSuite.BenchmarkPut1to3   5000000               837 ns/op
ArrayBlockingQueueSuite.BenchmarkPut1to4   1000000              1126 ns/op
ArrayBlockingQueueSuite.BenchmarkPut2to1   5000000               476 ns/op
ArrayBlockingQueueSuite.BenchmarkPut2to3   2000000               799 ns/op
ArrayBlockingQueueSuite.BenchmarkPut3to2   2000000               816 ns/op
ArrayBlockingQueueSuite.BenchmarkPut4to1   1000000              1239 ns/op
```
Having a different ratio of readers and writers introduce the same amount of latency.

## LICENCE
Copyright © 2017 Theo Despoudis MIT license
