package types

import "testing"

func TestIncUsesAPointerReceiver(t *testing.T) {
	var c Counter
	c.Inc()
	c.Inc()
	if c.Value() != 2 {
		t.Errorf("Value() = %d, want 2", c.Value())
	}
}

func TestIncByValueChangesNothing(t *testing.T) {
	var c Counter
	c.IncByValue()
	if c.Value() != 0 {
		t.Errorf("Value() = %d, want 0 - a value receiver gets a copy", c.Value())
	}
}

func TestMemStoreSatisfiesStore(t *testing.T) {
	var s Store = NewMemStore() // fails to compile if the methods are wrong

	s.Put("demo1", "web-1")

	got, ok := s.Get("demo1")
	if !ok || got != "web-1" {
		t.Errorf("Get() = %q, %v; want web-1, true", got, ok)
	}

	if _, ok := s.Get("demo9"); ok {
		t.Error("Get(demo9) reported present, want absent")
	}
}

func TestCopyAll(t *testing.T) {
	dst := NewMemStore()
	CopyAll(dst, map[string]string{"a": "1", "b": "2"})

	for k, want := range map[string]string{"a": "1", "b": "2"} {
		if got, ok := dst.Get(k); !ok || got != want {
			t.Errorf("Get(%q) = %q, %v; want %q, true", k, got, ok, want)
		}
	}
}
