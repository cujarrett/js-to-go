package values

import "testing"

func TestSetNote(t *testing.T) {
	var s Server
	SetNote(&s, "scheduled reboot")

	if s.Note == nil {
		t.Fatal("Note is nil, want a pointer to the note")
	}
	if *s.Note != "scheduled reboot" {
		t.Errorf("*Note = %q, want scheduled reboot", *s.Note)
	}
}

func TestNoteText(t *testing.T) {
	note := "scheduled reboot"
	tests := []struct {
		name   string
		server Server
		want   string
	}{
		{"note set", Server{Note: &note}, "scheduled reboot"},
		{"note nil", Server{}, ""},
		{"note empty string", Server{Note: new(string)}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoteText(tt.server); got != tt.want {
				t.Errorf("NoteText() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSlotOf(t *testing.T) {
	if got := SlotOf(&Server{Slot: "b2"}); got != "b2" {
		t.Errorf("SlotOf() = %q, want b2", got)
	}
	if got := SlotOf(nil); got != "" {
		t.Errorf("SlotOf(nil) = %q, want empty string", got)
	}
}
