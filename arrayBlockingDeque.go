package blockingQueues

import "sync"

type ArrayBlockingDeque struct {
	// Position of the first item of the Deque.
	first uint64

	// Position of the last item of the Deque.
	last uint64

	// The number of items in the Deque
	count uint64

	// Maximum number of items in the Deque
	capacity int

	// Main lock guarding all access
	lock *sync.RWMutex

	// Condition for waiting reads
	notEmpty *sync.Cond

	// Condition for waiting writes
	notFull *sync.Cond

	// The underling store
	store []interface{}
}
