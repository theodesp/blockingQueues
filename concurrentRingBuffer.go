package blockingQueues

import (
	"runtime"
	"sync/atomic"
)

type ConcurrentRingBuffer struct {
	// The padding members
	// below are here to ensure each item is on a separate cache line.
	pad1               [8]uint64
	lastCommittedIndex uint64
	pad2               [8]uint64
	writeIndex         uint64
	pad3               [8]uint64
	readIndex          uint64
	pad4               [8]uint64
	store              []interface{} // This will gain speed if its a specific type
	pad5               [8]uint64
}

func NewConcurrentRingBuffer(capacity uint64) *ConcurrentRingBuffer {
	return &ConcurrentRingBuffer{
		lastCommittedIndex: 0,
		writeIndex:         1,
		readIndex:          1,
		store:              make([]interface{}, capacity),
	}
}

func (q *ConcurrentRingBuffer) Put(value interface{}) (bool, error) {
	// Load next write index
	var nextWriteIndex = atomic.AddUint64(&q.writeIndex, 1) - 1
	var mask = uint64(cap(q.store) - 1)

	// Wait for reader to catch up as we don't want to go too far on the writes
	for nextWriteIndex > (q.readIndex + mask - 1) {
		// This will block the writer if the store is full
		runtime.Gosched()
	}

	// Write the item into it's slot
	q.store[nextWriteIndex&mask] = value

	// Increment the lastCommittedIndex so the item is available for reading
	for !atomic.CompareAndSwapUint64(&q.lastCommittedIndex, nextWriteIndex-1, nextWriteIndex) {
		runtime.Gosched()
	}

	return true, nil
}

func (q *ConcurrentRingBuffer) Get() (interface{}, error) {
	// Load next read index
	var nextReadIndex = atomic.AddUint64(&q.readIndex, 1) - 1
	var mask = uint64(cap(q.store) - 1)

	// If reader has out-run writer, wait for a value to be committed
	for nextReadIndex > q.lastCommittedIndex {
		runtime.Gosched()
	}
	return q.store[nextReadIndex&mask], nil
}
