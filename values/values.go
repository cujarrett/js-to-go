package values

// Server is the running state of one machine.
type Server struct {
	Name   string
	Active bool
	Note   *string // nil means "not set", which is different from ""
}

// Rename sets the server's name. The caller must see the change.
func Rename(s *Server, name string) {
	s.Name = name
}

// RenameCopy takes a Server by value. Set the name here too, then read the test
// to see why the caller is unaffected.
func RenameCopy(s Server, name string) {
	s.Name = name
}

// Load fills out with a server named after id and Active true. This is the shape
// every Kubernetes client call uses: you own the box, the function fills it.
func Load(id string, out *Server) error {
	out.Name = id
	out.Active = true
	return nil
}

// NoteOr returns the server's note, or fallback when no note is set.
func NoteOr(s Server, fallback string) string {
	if s.Note != nil {
		return *s.Note
	}
	return fallback
}
