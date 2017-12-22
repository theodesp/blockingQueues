package blockingQueues

import (
	"math"
	"sync"
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

	// store index for next write
	writeIndex uint64

	// store index for next read or remove
	readIndex uint64

	// The underling store
	store []interface{}
}

// Creates an ArrayBlockingQueue with the given (fixed) capacity
// returns an error if the capacity is lesst than 1
func NewArrayBlockingQueue(capacity uint64) (*ArrayBlockingQueue, error) {
	if capacity < 1 {
		return nil, ErrorCapacity
	}

	lock := new(sync.RWMutex)

	return &ArrayBlockingQueue{
		lock:     lock,
		notEmpty: sync.NewCond(lock.RLocker()),
		notFull:  sync.NewCond(lock.RLocker()),
		count:    uint64(0),
		store:    make([]interface{}, capacity),
	}, nil
}

// Returns the next increment of idx. Circulates the index
func (q ArrayBlockingQueue) inc(idx uint64) uint64 {
	if idx >= math.MaxUint64 {
		panic("Overflow")
	}

	if 1+idx == uint64(len(q.store)) {
		return 0
	} else {
		return idx + 1
	}
}

// Size returns this current elements size, is concurrent safe
func (q ArrayBlockingQueue) Size() uint64 {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.count
}

// Capacity returns this current elements remaining capacity, is concurrent safe
func (q ArrayBlockingQueue) Capacity() uint64 {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return uint64(uint64(len(q.store)) - q.count)
}

// Push element at current write position, advances, and signals.
// Call only when holding lock.
func (q *ArrayBlockingQueue) push(item interface{}) {
	q.store[q.writeIndex] = item
	q.writeIndex = q.inc(q.writeIndex)
	q.count += 1
	q.notEmpty.Signal()
}

// Pops element at current read position, advances, and signals.
// Call only when holding lock.
func (q *ArrayBlockingQueue) pop() interface{} {
	var item = q.store[q.readIndex]
	q.store[q.readIndex] = nil
	q.readIndex = q.inc(q.readIndex)
	q.count -= 1
	q.notFull.Signal()

	return item
}

// Pushes the specified element at the tail of the queue.
// Does not block the current goroutine
func (q *ArrayBlockingQueue) Push(item interface{}) (bool, error) {
	if q.Offer(item) {
		return true, nil
	} else {
		return false, ErrorFull
	}
}

// Inserts the specified element at the tail of this queue if it is possible to
// do so immediately without exceeding the queue's capacity,
// returning true upon success and false if this queue is full.
// Does not block the current goroutine
func (q *ArrayBlockingQueue) Offer(item interface{}) bool {
	if item == nil {
		panic("Null item")
	}

	q.lock.RLock()
	defer q.lock.RUnlock()

	if q.count == uint64(len(q.store)) {
		return false
	} else {
		q.push(item)
		return true
	}
}

// Pops an element from the head of the queue.
// Does not block the current goroutine
func (q *ArrayBlockingQueue) Pop() (interface{}, error) {
	q.lock.RLock()
	defer q.lock.RUnlock()

	if q.count == 0 {
		// Case empty
		return false, nil
	} else {
		var item = q.pop()
		return item, nil
	}
}

// Just attempts to return the tail element of the queue
func (q ArrayBlockingQueue) Peek() interface{} {
	q.lock.RLock()
	defer q.lock.RUnlock()

	if q.count == 0 {
		// Case empty
		return nil
	} else {
		var item = q.store[q.readIndex]
		return item
	}
}

func (q ArrayBlockingQueue) IsEmpty() bool {
	return q.Size() == 0
}

// Clears all the queues elements, cleans up, signals waiters for queue is empty
func (q *ArrayBlockingQueue) Clear() {
	q.lock.RLock()
	defer q.lock.RUnlock()

	// Start from head up to the tail
	next := q.readIndex

	for i := uint64(0); i < q.count; i += 1 {
		q.store[next] = nil
		next = q.inc(next)
	}

	q.count = uint64(0)
	q.readIndex = uint64(0)
	q.writeIndex = uint64(0)
	q.notFull.Broadcast()
}

// Takes an element from the head of the queue.
// It blocks the current goroutine if the queue is Empty until notified
func (q *ArrayBlockingQueue) Get() (interface{}, error) {
	q.lock.RLock()
	defer q.lock.RUnlock()

	for q.count == 0 {
		// We wait here until the queue has an item
		q.notEmpty.Wait()
	}

	// Critical section after wait released and predicate is false
	var item, err = q.Pop()
	return item, err
}

// Puts an element to the tail of the queue.
// It blocks the current goroutine if the queue is Full until notified
func (q *ArrayBlockingQueue) Put(item interface{}) (bool, error) {
	if item == nil {
		panic("Null item")
	}

	q.lock.RLock()
	defer q.lock.RUnlock()

	for q.count == uint64(len(q.store)) {
		// We wait here until the queue has an empty slot
		q.notFull.Wait()
	}

	// Critical section after wait released and predicate is false
	var res, err = q.Push(item)
	return res, err
}
