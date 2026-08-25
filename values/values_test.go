package values

import "testing"

func TestRenameChangesTheCaller(t *testing.T) {
	s := Server{Name: "original"}
	Rename(&s, "renamed")
	if s.Name != "renamed" {
		t.Errorf("Name = %q, want renamed", s.Name)
	}
}

func TestLoadFillsTheBox(t *testing.T) {
	var got Server
	if err := Load("web-1", &got); err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got.Name != "web-1" || !got.Active {
		t.Errorf("got %+v, want Name web-1 and Active true", got)
	}
}

func TestSetFirstReturnsTheChangedArray(t *testing.T) {
	ids := [3]string{"a", "b", "c"}
	got := SetFirst(ids, "z")

	if got != [3]string{"z", "b", "c"} {
		t.Errorf("got = %v, want [z b c]", got)
	}
	// ids itself is still {a, b, c} here - an array argument is a full copy, so
	// nothing SetFirst does could reach the caller's own variable. Not asserted:
	// true regardless of what SetFirst does, so it can never be the red test.
}

func TestNoteOr(t *testing.T) {
	note := "scheduled reboot"
	tests := []struct {
		name   string
		server Server
		want   string
	}{
		{"note set", Server{Note: &note}, "scheduled reboot"},
		{"note nil", Server{}, "none"},
		{"note empty string", Server{Note: new(string)}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoteOr(tt.server, "none"); got != tt.want {
				t.Errorf("NoteOr() = %q, want %q", got, tt.want)
			}
		})
	}
}
