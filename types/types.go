package types

// Counter counts things. The methods below show the difference a receiver makes.
type Counter struct {
	n int
}

// Inc adds one. The caller must see the change.
func (c *Counter) Inc() {
	// TODO
}

// IncByValue adds one to c, which is a copy of the caller's Counter.
// Write the increment anyway. Value() still reads 0 afterwards.
func (c Counter) IncByValue() {
	// TODO
}

// Value returns the current count.
func (c Counter) Value() int {
	// TODO
	return 0
}

// Store is anything that can hold and return a value by key.
// An interface is a list of methods. No type declares that it implements one.
type Store interface {
	Put(key, value string)
	Get(key string) (string, bool)
}

// MemStore keeps values in a map.
type MemStore struct {
	items map[string]string
}

// NewMemStore returns an empty MemStore ready to use.
func NewMemStore() *MemStore {
	// TODO: a nil map panics on write, so build one here
	return &MemStore{}
}

// Put stores a value.
func (m *MemStore) Put(key, value string) {
	// TODO
}

// Get returns a value and whether it was present.
func (m *MemStore) Get(key string) (string, bool) {
	// TODO
	return "", false
}

// CopyAll writes every pair into dst.
// dst is an interface, so any type with Put and Get fits. This is what
// "accept interfaces" means in practice.
func CopyAll(dst Store, pairs map[string]string) {
	// TODO
}
