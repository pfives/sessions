package sessions

// User represents one person who has access to the system.
type User interface {
	// GetID returns the user's unique ID.
	GetID() interface{}
}
