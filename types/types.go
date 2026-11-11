package types

// Counter counts things. The methods below show the difference a receiver makes.
type Counter struct {
	n int
}

// Inc adds one. The caller must see the change.
func (c *Counter) Inc() {
	c.n++
}

// IncByValue also adds one, but to a copy. Write it anyway, then read the test.
func (c Counter) IncByValue() {
	c.n++
}

// Value returns the current count.
func (c Counter) Value() int {
	return c.n
}

// Store is anything that can hold and return a value by key. An interface is a
// list of methods - no type ever declares that it implements one.
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
	return &MemStore{items: make(map[string]string)}
}

// Put stores a value.
func (m *MemStore) Put(key, value string) {
	m.items[key] = value
}

// Get returns a value and whether it was present.
func (m *MemStore) Get(key string) (string, bool) {
	v, ok := m.items[key]
	return v, ok
}

// CopyAll moves every pair from src into dst. It takes interfaces, so it works
// with any Store - this is what "accept interfaces" means in practice.
func CopyAll(dst Store, pairs map[string]string) {
	for k, v := range pairs {
		dst.Put(k, v)
	}
}
