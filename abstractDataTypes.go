package blockingQueues

type AbstractCollectionBase interface {
	Size() uint64
	Capacity() uint64
	IsEmpty() bool
	Clear()
}

// All Queues must implement this interface
type Interface interface {
	AbstractCollectionBase

	Push(item interface{}) (bool, error)
	Pop() (interface{}, error)

	Get() (interface{}, error)
	Put(item interface{}) (bool, error)
	Offer(item interface{}) bool

	Peek() interface{}
}

type QueueStore interface {
	Set(value interface{}, pos uint64)
	Remove(pos uint64) interface{}
	Get(pos uint64) interface{}
	Size() uint64
}
