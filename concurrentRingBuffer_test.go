package blockingQueues

import (
	. "gopkg.in/check.v1"
)

type ConcurrentRingBufferSuite struct {
	queue *ConcurrentRingBuffer
}

var _ = Suite(&ConcurrentRingBufferSuite{})

func (s *ConcurrentRingBufferSuite) SetUpTest(c *C) {
	s.queue = NewConcurrentRingBuffer(4096)
}

func (s *ConcurrentRingBufferSuite) BenchmarkRingBuffer1to1(c *C) {
	benchmarkPut(c, 1, 1, s.queue)
}

func (s *ConcurrentRingBufferSuite) BenchmarkRingBuffer2to2(c *C) {
	benchmarkPut(c, 2, 2, s.queue)
}

func (s *ConcurrentRingBufferSuite) BenchmarkRingBuffer4to4(c *C) {
	benchmarkPut(c, 4, 4, s.queue)
}

func (s *ConcurrentRingBufferSuite) BenchmarkRingBuffer4to1(c *C) {
	benchmarkPutMoreWriters(c, 4, 1, s.queue)
}

func (s *ConcurrentRingBufferSuite) BenchmarkRingBuffer2to1(c *C) {
	benchmarkPutMoreWriters(c, 2, 1, s.queue)
}

func (s *ConcurrentRingBufferSuite) BenchmarkRingBuffer3to2(c *C) {
	benchmarkPutMoreWriters(c, 3, 2, s.queue)
}

func (s *ConcurrentRingBufferSuite) BenchmarkRingBuffer2to3(c *C) {
	benchmarkPutMoreReaders(c, 2, 3, s.queue)
}

func (s *ConcurrentRingBufferSuite) BenchmarkRingBuffer1to3(c *C) {
	benchmarkPutMoreReaders(c, 1, 3, s.queue)
}

func (s *ConcurrentRingBufferSuite) BenchmarkRingBuffer1to4(c *C) {
	benchmarkPutMoreReaders(c, 1, 4, s.queue)
}
