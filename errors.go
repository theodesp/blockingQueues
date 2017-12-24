package blockingQueues

import "errors"

/**
 * Error definitions
 */

var ErrorCapacity = errors.New("ERROR_CAPACITY: attempt to Create Queue with invalid Capacity")
var ErrorFull = errors.New("ERROR_FULL: attempt to Put while Queue is Full")
var ErrorEmpty = errors.New("ERROR_EMPTY: attempt to Get while Queue is Empty")
