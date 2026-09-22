package values

// Second sitting. Everything here is about a pointer that might be nil, or a
// pointer you have to make yourself. Finish values.go first.

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
