package jsonv2

import (
	"strings"
	"testing"
	"time"
)

func TestEncodeOmitsZeroTime(t *testing.T) {
	out, err := Encode(Record{Name: "web-1"})
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	if strings.Contains(string(out), "lastSynced") {
		t.Errorf("got %s, want lastSynced omitted for a zero time.Time", out)
	}
}

func TestEncodeKeepsSetTime(t *testing.T) {
	when := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	out, err := Encode(Record{Name: "web-1", LastSynced: when})
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	if !strings.Contains(string(out), "lastSynced") {
		t.Errorf("got %s, want lastSynced present once set", out)
	}
}
