package observe

// Observer is used to moniter the transfer progress, caller should implement his own observer.
type Observer interface {
	// Transfer will be called when there are n bytes data tranfered.
	Transfer(n int)
	// Done will be called when transfer is done, or terminated by an error.
	Done(err error)
}
