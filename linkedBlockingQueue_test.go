package blockingQueues


import (
	"container/list"
	. "gopkg.in/check.v1"
	"math"
)

type LinkedBlockingQueueSuite struct {
	queue  *BlockingQueue
	queue2 *BlockingQueue
}

var _ = Suite(&LinkedBlockingQueueSuite{})

func (s *LinkedBlockingQueueSuite) SetUpTest(c *C) {
	s.queue, _ = NewLinkedBlockingQueue(16)
	s.queue2, _ = NewLinkedBlockingQueue(1024)
}

func (s *LinkedBlockingQueueSuite) TestInvalidCapacity(c *C) {
	_, err := NewLinkedBlockingQueue(0)
	c.Assert(err, ErrorMatches, "ERROR_CAPACITY: attempt to Create Queue with invalid Capacity")
}

func (s *LinkedBlockingQueueSuite) TestConstructor(c *C) {
	q, err := NewLinkedBlockingQueue(16)

	c.Assert(err, IsNil)
	c.Assert(q.Capacity(), Equals, uint64(16))
	c.Assert(q.Size(), Equals, uint64(0))
}

func (s *LinkedBlockingQueueSuite) TestIncrement(c *C) {
	defer func() {
		if r := recover(); r == nil {
			c.Errorf("TestIncrement should have panicked!")
		}
	}()

	c.Assert(s.queue.inc(0), Equals, uint64(1))
	c.Assert(s.queue.inc(16), Equals, uint64(17))
	c.Assert(s.queue.inc(20), Equals, uint64(21))
	c.Assert(s.queue.inc(math.MaxUint64), PanicMatches, "Overflow")

	s.queue.store.Set(1, uint64(0))
	s.queue.store.Set(2, uint64(1))
	c.Assert(s.queue.inc(2), Equals, uint64(0))
}

func (s *LinkedBlockingQueueSuite) TestPush(c *C) {
	for i := 0; i < 16; i += 1 {
		s.queue.Push(i)
	}

	c.Assert(s.queue.Size(), Equals, uint64(16))

	res, err := s.queue.Push(17)
	c.Assert(res, Equals, false)
	c.Assert(err, ErrorMatches, "ERROR_FULL: attempt to Put while Queue is Full")
}

func (s *LinkedBlockingQueueSuite) TestPop(c *C) {
	for i := 0; i < 10; i += 1 {
		s.queue.Push(i)
	}

	for i := 0; i < 10; i += 1 {
		s.queue.Pop()
	}

	c.Assert(s.queue.Size(), Equals, uint64(0))

	res, err := s.queue.Pop()
	c.Assert(res, IsNil)
	c.Assert(err, ErrorMatches, "ERROR_EMPTY: attempt to Get while Queue is Empty")
}

func (s *LinkedBlockingQueueSuite) TestClear(c *C) {
	for i := 0; i < 10; i += 1 {
		s.queue.Push(i)
	}

	s.queue.Clear()

	c.Assert(s.queue.Size(), Equals, uint64(0))
}

func (s *LinkedBlockingQueueSuite) TestPeek(c *C) {
	for i := 0; i < 10; i += 1 {
		s.queue.Push(i)
	}

	c.Assert(s.queue.Peek().(*list.Element).Value, Equals, 0)

	s.queue.Pop()

	c.Assert(s.queue.Peek().(*list.Element).Value, Equals, 1)
}

func (s *LinkedBlockingQueueSuite) TestPutPanicsOnNil(c *C) {
	defer func() {
		if r := recover(); r == nil {
			c.Errorf("TestPutNotFull should have panicked!")
		}
	}()

	s.queue.Put(nil)
}

func (s *LinkedBlockingQueueSuite) TestPutBlocks(c *C) {
	q := s.queue
	for i := 0; i < 16; i += 1 {
		q.Push(i)
	}

	c.Assert(q.Size(), Equals, uint64(16))

	n := 2
	running := make(chan bool, n)
	awake := make(chan bool, n)

	// The next 2 items will block
	for i := 0; i < n; i++ {
		go func(i int) {
			running <- true
			q.Put(i)
			awake <- true
		}(i)
	}
	for i := 0; i < n; i++ {
		<-running // Wait for everyone to run.
	}
	j := 0
	for n > 0 {
		select {
		case <-awake:
			c.Error("BlockingArray not asleep")
		default:
		}
		item, err := q.Get()
		c.Assert(err, IsNil)
		c.Assert(item, Equals, j)
		<-awake // Will deadlock if no goroutine wakes up
		select {
		case <-awake:
			c.Error("too many waiters awake")
		default:
		}
		n--
		j += 1
	}
	thirdItem, err2 := q.Get()

	c.Assert(err2, IsNil)
	c.Assert(thirdItem, Equals, 2)
}

func (s *LinkedBlockingQueueSuite) BenchmarkPeek(c *C) {
	for i := 0; i < c.N; i++ {
		s.queue.Peek()
	}

	q, _ := NewLinkedBlockingQueue(math.MaxUint16)

	q.Push(1)
	q.Push(1)
	q.Push(1)

	c.ResetTimer()

	for i := 0; i < c.N; i++ {
		s.queue.Peek()
	}
}

func (s *LinkedBlockingQueueSuite) BenchmarkPopOverflow(c *C) {
	for i := 0; i < c.N; i++ {
		s.queue.Pop()
	}
}

func (s *LinkedBlockingQueueSuite) BenchmarkPop(c *C) {
	q, _ := NewLinkedBlockingQueue(math.MaxUint16)

	for i := 0; i < c.N; i++ {
		q.Push(i)
	}

	c.ResetTimer()

	for i := 0; i < c.N; i++ {
		q.Pop()
	}
}

func (s *LinkedBlockingQueueSuite) BenchmarkPushOverflow(c *C) {
	for i := 0; i < c.N; i++ {
		s.queue.Push(i)
	}
}

func (s *LinkedBlockingQueueSuite) BenchmarkPush(c *C) {
	q, _ := NewLinkedBlockingQueue(math.MaxUint16)

	c.ResetTimer()

	for i := 0; i < c.N; i++ {
		q.Push(i)
	}
}

func (s *LinkedBlockingQueueSuite) BenchmarkPut1to1(c *C) {
	benchmarkPut(c, 1, 1, s.queue2)
}

func (s *LinkedBlockingQueueSuite) BenchmarkPut2to2(c *C) {
	benchmarkPut(c, 2, 2, s.queue2)
}

func (s *LinkedBlockingQueueSuite) BenchmarkPut4to4(c *C) {
	benchmarkPut(c, 4, 4, s.queue2)
}

func (s *LinkedBlockingQueueSuite) BenchmarkPut4to1(c *C) {
	benchmarkPutMoreWriters(c, 4, 1, s.queue2)
}

func (s *LinkedBlockingQueueSuite) BenchmarkPut2to1(c *C) {
	benchmarkPutMoreWriters(c, 2, 1, s.queue2)
}

func (s *LinkedBlockingQueueSuite) BenchmarkPut3to2(c *C) {
	benchmarkPutMoreWriters(c, 3, 2, s.queue2)
}

func (s *LinkedBlockingQueueSuite) BenchmarkPut1to3(c *C) {
	benchmarkPutMoreReaders(c, 1, 3, s.queue2)
}

func (s *LinkedBlockingQueueSuite) BenchmarkPut2to3(c *C) {
	benchmarkPutMoreReaders(c, 2, 3, s.queue2)
}

func (s *LinkedBlockingQueueSuite) BenchmarkPut1to4(c *C) {
	benchmarkPutMoreReaders(c, 1, 4, s.queue2)
}
