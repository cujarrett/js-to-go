package collections

// Server is one machine in an inventory.
type Server struct {
	Name   string
	Slot   string
	Active bool
}

// Names returns every server's name, in order.
func Names(servers []Server) []string {
	names := make([]string, 0, len(servers))
	for _, s := range servers {
		names = append(names, s.Name)
	}
	return names
}

// SlotSet returns a lookup of the slots in use. A map[string]bool is Go's Set.
func SlotSet(servers []Server) map[string]bool {
	slots := map[string]bool{}
	for _, s := range servers {
		slots[s.Slot] = true
	}
	return slots
}

// CountBySlot counts how many servers hold each slot. Note that reading a
// missing key gives the zero value rather than an error.
func CountBySlot(servers []Server) map[string]int {
	counts := map[string]int{}
	for _, s := range servers {
		counts[s.Slot]++
	}
	return counts
}

// ActivateAll sets Active on every server. Ranging gives you a copy of each
// element, so this needs the index.
func ActivateAll(servers []Server) {
	for i := range servers {
		servers[i].Active = true
	}
}

// Unselected returns the names of servers whose slot is not in keep.
func Unselected(servers []Server, keep map[string]bool) []string {
	var names []string
	for _, s := range servers {
		if !keep[s.Slot] {
			names = append(names, s.Name)
		}
	}
	return names
}

// AddServer appends s and returns the grown slice. append does not grow the
// caller's slice in place - it builds a new one and hands it back, so a call
// that discards the return value changes nothing the caller can see.
func AddServer(servers []Server, s Server) []Server {
	return append(servers, s)
}
