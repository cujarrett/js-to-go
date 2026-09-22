package values

import "testing"

func TestActivate_ChangesTheCaller(t *testing.T) {
	s := Server{Name: "web-1"}
	Activate(&s)
	if !s.Active {
		t.Error("Active = false, want true")
	}
}

func TestFetch_FillsTheCallersServer(t *testing.T) {
	var got Server
	if err := Fetch("web-1", &got); err != nil {
		t.Fatalf("Fetch: %v", err)
	}
	if got.Name != "web-1" || got.Slot != "a1" || !got.Active {
		t.Errorf("got %+v, want Name web-1, Slot a1, Active true", got)
	}
}

func TestSetLast_ReturnsTheChangedArray(t *testing.T) {
	slots := [3]string{"a", "b", "c"}
	got := SetLast(slots, "z")

	if got != [3]string{"a", "b", "z"} {
		t.Errorf("got = %v, want [a b z]", got)
	}
	// slots itself is still {a, b, c} here. An array argument is a full copy, so
	// nothing SetLast does could reach the caller's own variable. Not asserted:
	// true regardless of what SetLast does, so it can never be the red test.
}
