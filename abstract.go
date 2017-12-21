package blockingQueues

type AbstractCollectionBase interface {
	Size() uint64
	Capacity() uint64
	IsEmpty() bool
	Clear()
}

// All Queues must implement this interface
type Interface interface {
	Push(item interface{}) (bool, error)
	Pop() (interface{}, error)

	Poll() (interface{}, error)
	Offer(item interface{}) bool

	Peek() interface{}

	AbstractCollectionBase
}
