package observe

// Caller should implement his own Observer to moniter the transfer process.
// All methods on Observer will be called in blocking mode, caller should
// implement them carefully.
type Observer interface {
	// Transfer will be called when there are n bytes data tranfered.
	Transfer(n int)
	// Done will be called when transfer is done, or terminated by an error.
	Done(err error)
}
