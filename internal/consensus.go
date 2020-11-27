package internal

// Explain which kind of operation is applied to the consensus
// algorithm.
type Operation uint8

const (
	// Operation to tick the clock value.
	Tick Operation = iota

	// Operation to leap the clock value.
	Leap
)

// Structure to hold commands issued to the consensus protocol.
type Command struct {
	// Defines the operation that will be applied.
	Operation Operation `json:"operation,omitempty"`

	// Carry the value that will be applied.
	// This is only used in cases where the Operation is
	// of type Leap.
	Value uint64 `json:"value,omitempty"`
}

// Consensus interface to abstract the current protocol.
type Consensus interface {
	// Propose a new command to the protocol.
	Propose(Command) error

	// Read the current value available.
	Read() (uint64, error)

	// Close and finishes the consensus algorithm,
	// any commands or reads issued after will panic.
	Close() error
}
