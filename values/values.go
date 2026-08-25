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

// RenameCopy takes a Server by value. Set the name here too, then read the test
// to see why the caller is unaffected.
func RenameCopy(s Server, name string) {
	// TODO
}

// Load fills out with a server named after id and Active true. This is the shape
// every Kubernetes client call uses: you own the box, the function fills it.
func Load(id string, out *Server) error {
	// TODO
	return nil
}

// NoteOr returns the server's note, or fallback when no note is set.
func NoteOr(s Server, fallback string) string {
	// TODO
	return ""
}

// SetFirst sets index 0 of ids and returns the changed array. An array is data,
// copied whole when passed - the same as Server above, with no address inside it
// for a change to travel back through. That is why this returns, like RenameCopy
// would have to, rather than mutating in place.
func SetFirst(ids [3]string, id string) [3]string {
	// TODO
	return ids
}
