package values

// Server is the running state of one machine.
type Server struct {
	Name   string
	Active bool
	Note   *string // nil means "not set", which is different from ""
}

// Rename sets the server's name. The caller must see the change.
func Rename(s *Server, name string) {
	// TODO
}

// RenameCopy takes a Server by value, so s is a copy of the caller's Server.
// Set the name on it anyway. There is no address back to the original, so
// nothing written here can reach the caller. That is the whole point of it.
func RenameCopy(s Server, name string) {
	// TODO
}

// Load writes a server named id, with Active true, into out.
// The caller declares the Server and passes its address; Load fills it in.
// Every Kubernetes client Get works this way.
func Load(id string, out *Server) error {
	// TODO
	return nil
}

// NoteOr returns the server's note, or fallback when no note is set.
func NoteOr(s Server, fallback string) string {
	// TODO
	return ""
}

// SetFirst sets ids[0] to id and returns the array.
// A [3]string is data like Server above, so the parameter is a full copy.
// Returning it is the only way the caller can see the change.
func SetFirst(ids [3]string, id string) [3]string {
	// TODO
	return ids
}
