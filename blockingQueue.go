package blockingQueues

import (
	"math"
	"sync"
)

/**
 * BlockingQueue is A multi-producer, multi-consumer queue
 */

type BlockingQueue struct {
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
	store QueueStore
}

// Returns the next increment of idx. Circulates the index
func (q BlockingQueue) inc(idx uint64) uint64 {
	if idx >= math.MaxUint64 {
		panic("Overflow")
	}

	if 1+idx == q.store.Size() {
		return 0
	} else {
		return idx + 1
	}
}

// Size returns this current elements size, is concurrent safe
func (q BlockingQueue) Size() uint64 {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.count
}

// Capacity returns this current elements remaining capacity, is concurrent safe
func (q BlockingQueue) Capacity() uint64 {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return uint64(q.store.Size() - q.count)
}

// Push element at current write position, advances, and signals.
// Call only when holding lock.
func (q *BlockingQueue) push(item interface{}) {
	q.store.Set(item, q.writeIndex)
	q.writeIndex = q.inc(q.writeIndex)
	q.count += 1
	q.notEmpty.Signal()
}

// Pops element at current read position, advances, and signals.
// Call only when holding lock.
func (q *BlockingQueue) pop() interface{} {
	var item = q.store.Remove(q.readIndex)
	q.readIndex = q.inc(q.readIndex)
	q.count -= 1
	q.notFull.Signal()

	return item
}

// Pushes the specified element at the tail of the queue.
// Does not block the current goroutine
func (q *BlockingQueue) Push(item interface{}) (bool, error) {
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
func (q *BlockingQueue) Offer(item interface{}) bool {
	if item == nil {
		panic("Null item")
	}

	q.lock.RLock()
	defer q.lock.RUnlock()

	if q.count == q.store.Size() {
		return false
	} else {
		q.push(item)
		return true
	}
}

// Pops an element from the head of the queue.
// Does not block the current goroutine
func (q *BlockingQueue) Pop() (interface{}, error) {
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
func (q BlockingQueue) Peek() interface{} {
	q.lock.RLock()
	defer q.lock.RUnlock()

	if q.count == 0 {
		// Case empty
		return nil
	} else {
		var item = q.store.Get(q.readIndex)
		return item
	}
}

func (q BlockingQueue) IsEmpty() bool {
	return q.Size() == 0
}

// Clears all the queues elements, cleans up, signals waiters for queue is empty
func (q *BlockingQueue) Clear() {
	q.lock.RLock()
	defer q.lock.RUnlock()

	// Start from head up to the tail
	next := q.readIndex

	for i := uint64(0); i < q.count; i += 1 {
		q.store.Set(nil, next)
		next = q.inc(next)
	}

	q.count = uint64(0)
	q.readIndex = uint64(0)
	q.writeIndex = uint64(0)
	q.notFull.Broadcast()
}

// Takes an element from the head of the queue.
// It blocks the current goroutine if the queue is Empty until notified
func (q *BlockingQueue) Get() (interface{}, error) {
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
func (q *BlockingQueue) Put(item interface{}) (bool, error) {
	if item == nil {
		panic("Null item")
	}

	q.lock.RLock()
	defer q.lock.RUnlock()

	for q.count == q.store.Size() {
		// We wait here until the queue has an empty slot
		q.notFull.Wait()
	}

	// Critical section after wait released and predicate is false
	var res, err = q.Push(item)
	return res, err
}
