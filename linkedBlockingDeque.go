package blockingQueues

import (
	"container/list"
	"sync"
)

type LinkedBlockingDeque struct {
	// Maximum number of items in the Deque
	capacity int

	// Main lock guarding all access
	lock *sync.RWMutex

	// Condition for waiting reads
	notEmpty *sync.Cond

	// Condition for waiting writes
	notFull *sync.Cond

	// The underling store
	store list.List
}
