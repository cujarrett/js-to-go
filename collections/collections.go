package collections

// Server is one machine in an inventory.
type Server struct {
	Name   string
	Slot   string
	Active bool
}

// Names returns every server's name, in order.
func Names(servers []Server) []string {
	result := []string{}
	for _, s := range servers {
		result = append(result, s.Name)
	}
	return result
}

// SlotSet returns a lookup of the slots in use. A map[string]bool is Go's Set.
func SlotSet(servers []Server) map[string]bool {
	result := map[string]bool{}
	for _, s := range servers {
		result[s.Slot] = true
	}
	return result
}

// CountBySlot counts how many servers hold each slot. Note that reading a
// missing key gives the zero value rather than an error.
func CountBySlot(servers []Server) map[string]int {
	result := map[string]int{}
	for _, s := range servers {
		result[s.Slot] += 1
	}
	return result
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
	result := []string{}

	for _, s := range servers {
		if !keep[s.Slot] {
			result = append(result, s.Name)
		}
	}
	return result
}
