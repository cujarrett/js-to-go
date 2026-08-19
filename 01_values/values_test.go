package values

import "testing"

func TestRenameChangesTheCaller(t *testing.T) {
	s := Server{Name: "original"}
	Rename(&s, "renamed")
	if s.Name != "renamed" {
		t.Errorf("Name = %q, want renamed", s.Name)
	}
}

func TestRenameCopyDoesNot(t *testing.T) {
	s := Server{Name: "original"}
	RenameCopy(s, "renamed")
	if s.Name != "original" {
		t.Errorf("Name = %q, want original - a copy cannot change the caller", s.Name)
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
