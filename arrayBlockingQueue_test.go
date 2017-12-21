package blockingQueues

import (
	"testing"

	. "gopkg.in/check.v1"
	"math"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ArrayBlockingQueueSuite struct {
	queue *ArrayBlockingQueue
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

	s.queue.store = append(s.queue.store, 1)
	s.queue.store = append(s.queue.store, 2)
	c.Assert(s.queue.inc(2), Equals, uint64(0))
}
