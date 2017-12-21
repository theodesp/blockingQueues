package blockingQueues

import (
	"sync"
	"math"
)

/**
 * ArrayBlockingQueue is A multi-producer, multi-consumer queue
 */

type ArrayBlockingQueue struct {
	// The number of items in the Queue
	count uint64

	// Main lock guarding all access
	lock *sync.RWMutex

	// Condition for waiting reads
	notEmpty *sync.Cond

	// Condition for waiting writes
	notFull *sync.Cond

	// The underling store
	store []interface{}
}

// Creates an ArrayBlockingQueue with the given (fixed) capacity
// returns an error if the capacity is lesst than 1
func NewArrayBlockingQueue(capacity uint64) (*ArrayBlockingQueue, error)  {
	if capacity < 1 {
		return nil, ErrorCapacity
	}

	lock := new(sync.RWMutex)

	return &ArrayBlockingQueue{
		lock: lock,
		notEmpty: sync.NewCond(lock.RLocker()),
		notFull: sync.NewCond(lock.RLocker()),
		count: uint64(0),
		store: make([]interface{}, capacity),
	}, nil
}


// Returns the next increment of idx. Circulates the index
func (q ArrayBlockingQueue) inc(idx uint64) uint64 {
	if idx >= math.MaxUint64 {
		panic("Overflow")
	}

	var i = idx + 1
	if i == uint64(len(q.store)) {
		return 0
	} else {
		return i
	}
}

// Size returns this current elements size, is concurrent safe
func (q ArrayBlockingQueue) Size() uint64 {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.count
}

// Capacity returns this current elements remaining capacity, is concurrent safe
func (q ArrayBlockingQueue) Capacity() uint64 {
	q.lock.Lock()
	defer q.lock.Unlock()

	return uint64(uint64(len(q.store)) - q.count)
}
