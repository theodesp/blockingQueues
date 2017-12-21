package blockingQueues

// All Queues must implement this minimal interface
type Interface interface {
	Offer(value interface{}) (bool, error)
	Poll() (interface{}, error)
	Peek() (interface{}, error)
	Size() uint64
	Capacity() uint64
	Empty() bool
	Clear()
	Includes(value interface{}) bool
	Remove(index uint64) (bool, error)
}
