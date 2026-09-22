package collections

// Server is one machine in an inventory.
type Server struct {
	Name   string
	Slot   string
	Active bool
}

// Names returns every server's name, in order.
func Names(servers []Server) []string {
	// TODO
	return nil
}

// SlotSet returns a lookup of the slots in use. A map[string]bool is Go's Set.
func SlotSet(servers []Server) map[string]bool {
	// TODO
	return nil
}

// CountBySlot counts how many servers hold each slot. Note that reading a
// missing key gives the zero value rather than an error.
func CountBySlot(servers []Server) map[string]int {
	// TODO
	return nil
}

// ActivateAll sets Active on every server. Ranging gives you a copy of each
// element, so this needs the index.
func ActivateAll(servers []Server) {
	// TODO
}

// Unselected returns the names of servers whose slot is not in keep.
func Unselected(servers []Server, keep map[string]bool) []string {
	// TODO
	return nil
}

// AddServer appends s and returns the grown slice. append does not grow the
// caller's slice in place - it builds a new one and hands it back, so a call
// that discards the return value changes nothing the caller can see.
func AddServer(servers []Server, s Server) []Server {
	// TODO
	return nil
}
