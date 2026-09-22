package values

// Server is the running state of one machine.
type Server struct {
	Name   string
	Slot   string
	Active bool
	Note   *string // nil means "not set", which is different from ""
}

// Activate marks the server active. The caller must see the change.
func Activate(s *Server) {
	// TODO
}

// ActivateCopy takes a Server by value, so s is a copy of the caller's Server.
// Set Active on it anyway. There is no address back to the original, so nothing
// written here can reach the caller. That is the whole point of it.
func ActivateCopy(s Server) {
	// TODO
}

// Fetch writes a server named id, in slot "a1" and active, into out.
// The caller declares the Server and passes its address; Fetch fills it in.
// Every Kubernetes client Get works this way.
func Fetch(id string, out *Server) error {
	// TODO
	return nil
}

// SetNote points the server's Note at note.
// Note is a *string, so a plain assignment will not compile. You need an address.
// note is a parameter, so it is already this function's own copy of the caller's
// string, and taking its address cannot alias anything the caller still holds.
func SetNote(s *Server, note string) {
	// TODO
}

// NoteText returns the server's note, or "" when no note is set.
// Dereferencing a nil pointer panics, so check before reading.
func NoteText(s Server) string {
	// TODO
	return ""
}

// SlotOf returns the server's slot, or "" when s is nil.
// A *Server parameter can arrive nil. Reading a field off it panics.
func SlotOf(s *Server) string {
	// TODO
	return ""
}

// SetLast sets the final element of slots to s and returns the array.
// A [3]string is data like Server above, so the parameter is a full copy.
// Returning it is the only way the caller can see the change.
func SetLast(slots [3]string, s string) [3]string {
	// TODO
	return slots
}
