package blockingQueues

import (
	"testing"

	. "gopkg.in/check.v1"
	"math"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ArrayBlockingQueueSuite struct {
	queue *BlockingQueue
}

var _ = Suite(&ArrayBlockingQueueSuite{})

func (s *ArrayBlockingQueueSuite) SetUpTest(c *C) {
	s.queue, _ = NewArrayBlockingQueue(16)
}

func (s *ArrayBlockingQueueSuite) TestInvalidCapacity(c *C) {
	_, err := NewArrayBlockingQueue(0)
	c.Assert(err, ErrorMatches, "ERROR_CAPACITY: attempt to Create Queue with invalid Capacity")
}

func (s *ArrayBlockingQueueSuite) TestContstructor(c *C) {
	q, err := NewArrayBlockingQueue(16)

	c.Assert(err, IsNil)
	c.Assert(q.Capacity(), Equals, uint64(16))
	c.Assert(q.Size(), Equals, uint64(0))
}

func (s *ArrayBlockingQueueSuite) TestIncrement(c *C) {
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

func (s *ArrayBlockingQueueSuite) TestPush(c *C) {
	for i:= 0;i < 10; i+=1 {
		s.queue.Push(i)
	}

	c.Assert(s.queue.Size(), Equals, uint64(10))
}


func (s *ArrayBlockingQueueSuite) BenchmarkPushOverflow(c *C) {
	for i := 0; i < c.N; i++ {
		s.queue.Push(i)
	}
}

func (s *ArrayBlockingQueueSuite) BenchmarkPush(c *C) {
	q, _ := NewArrayBlockingQueue(math.MaxUint16)

	c.ResetTimer()

	for i := 0; i < c.N; i++ {
		q.Push(i)
	}
}

func (s *ArrayBlockingQueueSuite) TestPop(c *C) {
	for i:= 0;i < 10; i+=1 {
		s.queue.Push(i)
	}

	for i:= 0;i < 10; i+=1 {
		s.queue.Pop()
	}

	c.Assert(s.queue.Size(), Equals, uint64(0))
}


func (s *ArrayBlockingQueueSuite) BenchmarkPopOverflow(c *C) {
	for i := 0; i < c.N; i++ {
		s.queue.Pop()
	}
}

func (s *ArrayBlockingQueueSuite) BenchmarkPop(c *C) {
	q, _ := NewArrayBlockingQueue(math.MaxUint16)

	for i := 0; i < c.N; i++ {
		q.Push(i)
	}

	c.ResetTimer()
	c.StartTimer()

	for i := 0; i < c.N; i++ {
		q.Pop()
	}
}