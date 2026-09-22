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

// SetLast sets the final element of slots to s and returns the array.
// A [3]string is data like Server above, so the parameter is a full copy.
// Returning it is the only way the caller can see the change.
func SetLast(slots [3]string, s string) [3]string {
	// TODO
	return slots
}
