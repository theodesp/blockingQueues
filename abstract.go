package blockingQueues

// All Queues must implement this interface
type Interface interface {
	Push(value interface{}) (bool, error)
	Pop()(interface{}, error)
	Poll() (interface{}, error)
	Peek() interface{}
	Size() uint64
	Capacity() uint64
	Empty() bool
	Clear()
	Includes(value interface{}) bool
	Remove(index uint64) (bool, error)
}
