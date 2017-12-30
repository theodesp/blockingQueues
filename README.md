Blocking Queues
---
<a href="https://godoc.org/github.com/theodesp/blockingQueues">
<img src="https://godoc.org/github.com/theodesp/blockingQueues?status.svg" alt="GoDoc">
</a>

<a href="https://opensource.org/licenses/MIT" rel="nofollow">
<img src="https://img.shields.io/github/license/mashape/apistatus.svg" alt="License"/>
</a>

<a href="https://travis-ci.org/theodesp/blockingQueues" rel="nofollow">
<img src="https://travis-ci.org/theodesp/blockingQueues.svg?branch=master" />
</a>

<a href="https://codecov.io/gh/theodesp/blockingQueues">
  <img src="https://codecov.io/gh/theodesp/blockingQueues/branch/master/graph/badge.svg" />
</a>

<a href="https://ci.appveyor.com/project/theodesp/blockingqueues" rel="nofollow">
  <img src="https://ci.appveyor.com/api/projects/status/7yiwtn68qmcj71xy?svg=true" />
</a>

Blocking Queues provides some simple, performant, goroutine safe queues useful as resource pools or job queues. 
The primary focus is simplicity and high performance without sacrificing readability. In fact I tried to
provide good documentation on the code and some examples of usage.


## Queues Provided
* **ArrayBlockingQueue**: A bounded blocking queue backed by a slice
* **LinkedBlockingQueue**: A bounded blocking queue backed by a container/list
* **ConcurrentRingBuffer**: A bounded lock-free queue backed by a slice

## Installation
```go
go get -u github.com/theodesp/blockingQueues
```

## Usage

Non blocking api
```go
queue, _ := NewArrayBlockingQueue(2)
res, _ := queue.Push(1)
res, _ := queue.Push(2)
res, err := queue.Push(3) // err is not Nil as queue is full
res, err := queue.Pop()
res, err := queue.Pop()
res, err := queue.Pop() // err is not Nil as queue is empty
```

Blocking api
```go
queue, _ := NewArrayBlockingQueue(2)
res, _ := queue.Put(1)
res, _ := queue.Put(2)
res, err := queue.Put(3) // Will block the current goroutine

// In another goroutine
res, err := queue.Get() // Will unblock the first goroutine and add the last item
res, err := queue.Get()
res, err := queue.Get()
res, err := queue.Get() // Will block the current goroutine
```

Full API Documentation: 
[https://godoc.org/github.com/theodesp/blockingQueues](https://godoc.org/github.com/theodesp/blockingQueues)

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
#### ArrayBlockingQueue
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

#### LinkedBlockingQueue
```text
LinkedBlockingQueueSuite.BenchmarkPeek       100000000               21.4 ns/op
LinkedBlockingQueueSuite.BenchmarkPop100000000               24.4 ns/op
LinkedBlockingQueueSuite.BenchmarkPopOverflow        100000000               23.4 ns/op
LinkedBlockingQueueSuite.BenchmarkPush       50000000                47.3 ns/op
LinkedBlockingQueueSuite.BenchmarkPushOverflow       50000000                42.1 ns/op
LinkedBlockingQueueSuite.BenchmarkPut1to1    10000000               246 ns/op
LinkedBlockingQueueSuite.BenchmarkPut1to3     2000000               930 ns/op
LinkedBlockingQueueSuite.BenchmarkPut1to4     1000000              1496 ns/op
LinkedBlockingQueueSuite.BenchmarkPut2to1     5000000               578 ns/op
LinkedBlockingQueueSuite.BenchmarkPut2to2     5000000               560 ns/op
LinkedBlockingQueueSuite.BenchmarkPut2to3     2000000              1053 ns/op
LinkedBlockingQueueSuite.BenchmarkPut3to2     2000000              1041 ns/op
LinkedBlockingQueueSuite.BenchmarkPut4to1     1000000              1488 ns/op
LinkedBlockingQueueSuite.BenchmarkPut4to4     1000000              1451 ns/op
```

#### ConcurrentRingBuffer
Test
```text
ConcurrentRingBufferSuite.BenchmarkRingBuffer1to1        20000000                85.7 ns/op
ConcurrentRingBufferSuite.BenchmarkRingBuffer1to3         1000000              2793 ns/op
ConcurrentRingBufferSuite.BenchmarkRingBuffer1to4          500000              5501 ns/op
ConcurrentRingBufferSuite.BenchmarkRingBuffer2to1         5000000               465 ns/op
ConcurrentRingBufferSuite.BenchmarkRingBuffer2to2         5000000               474 ns/op
ConcurrentRingBufferSuite.BenchmarkRingBuffer2to3         1000000              2640 ns/op
ConcurrentRingBufferSuite.BenchmarkRingBuffer3to2         1000000              2766 ns/op
ConcurrentRingBufferSuite.BenchmarkRingBuffer4to1         1000000              5411 ns/op
ConcurrentRingBufferSuite.BenchmarkRingBuffer4to4          500000              5370 ns/op

```

## LICENCE
Copyright © 2017 Theo Despoudis MIT license
