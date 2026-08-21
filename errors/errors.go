package errors

import "errors"

// ErrNotFound is a sentinel: a specific error value callers can test for.
var ErrNotFound = errors.New("not found")

// Store is a tiny in-memory lookup.
type Store struct {
	items map[string]string
}

// NewStore builds a Store from pairs.
func NewStore(items map[string]string) *Store {
	return &Store{items: items}
}

// Get returns the value for key, or ErrNotFound when it is absent.
func (s *Store) Get(key string) (string, error) {
	// TODO
	return "", nil
}

// Describe returns "key=value" for key. When Get fails, wrap the error with
// context using %w so the caller can still recognise ErrNotFound.
func (s *Store) Describe(key string) (string, error) {
	// TODO
	return "", nil
}

// MissingKeyError says which key was missing, for callers that want the detail
// rather than just the fact.
type MissingKeyError struct {
	Key string
}

// Error makes MissingKeyError satisfy the error interface.
func (e *MissingKeyError) Error() string {
	// TODO: return something like `key "demo9" is missing`
	return ""
}

// Require returns a *MissingKeyError when key is absent, and nil otherwise.
func (s *Store) Require(key string) error {
	// TODO
	return nil
}
