// Package generics is what typed clients and generic informers in client-go
// and controller-runtime are built on: a function or type that works over
// many types, checked at compile time, without falling back to any.
package generics

// Names returns the Name field of every item - the generic cousin of module
// 02's Names, which only ever worked on []Server.
//
// T any means "T can be anything." Named lets the compiler prove every T has
// a Name field, without T being one fixed type.
func Names[T Named](items []T) []string {
	// TODO: same loop as module 02's Names, calling item.Name() instead of item.Name
	return nil
}

// Named is a constraint: not "a list of methods a value has", the way module
// 04's Store was, but "a list of methods a type parameter must satisfy to be
// used here." Constraints and interfaces share the same syntax on purpose.
type Named interface {
	Name() string
}

// Server satisfies Named, so []Server can be passed to Names[Server].
type Server struct {
	name string
}

func (s Server) Name() string { return s.name }

// Pod also satisfies Named - a different type, same constraint, same Names
// function. Nothing about Names changed to accept it.
type Pod struct {
	podName string
}

func (p Pod) Name() string { return p.podName }

// Set is a generic version of module 02's map[string]bool trick - a set of any
// comparable type, not just string.
type Set[T comparable] struct {
	items map[T]bool
}

// NewSet returns an empty Set.
func NewSet[T comparable]() *Set[T] {
	// TODO
	return nil
}

// Add puts v in the set.
func (s *Set[T]) Add(v T) {
	// TODO
}

// Has reports whether v is in the set.
func (s *Set[T]) Has(v T) bool {
	// TODO
	return false
}
