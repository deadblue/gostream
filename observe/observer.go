package observe

// Caller need implement his own Observer to moniter the transfer process.
// All methods on Observer will be called in blocking mode, caller should
// take care in their implementation.
type Observer interface {
	// OnTransfer will be called when there are n bytes data tranfered.
	OnTransfer(n int)
	// OnStop will be called when transfer is done, or terminated by an error.
	OnStop(err error)
}
