package json

import (
	"strings"
	"testing"
)

func TestEncodeUsesTagNames(t *testing.T) {
	out, err := Encode(Config{SourceSecret: "demo1-tls", Targets: []string{"demo1"}})
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}

	got := string(out)
	for _, want := range []string{`"sourceSecret":"demo1-tls"`, `"targets":["demo1"]`} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %s in %s", want, got)
		}
	}
}

func TestEncodeOmitsEmptyOptionalFields(t *testing.T) {
	out, _ := Encode(Config{SourceSecret: "demo1-tls", Targets: []string{"demo1"}})

	got := string(out)
	if strings.Contains(got, "replicas") || strings.Contains(got, "note") {
		t.Errorf("got %s, want replicas and note omitted when empty", got)
	}
	if !strings.Contains(got, "targets") {
		t.Errorf("got %s, want targets present even so", got)
	}
}

func TestEncodeNeverIncludesUnexportedFields(t *testing.T) {
	out, _ := Encode(Config{SourceSecret: "demo1-tls", internal: "hidden"})
	if strings.Contains(string(out), "hidden") {
		t.Errorf("got %s, want no unexported field", out)
	}
}

func TestDecode(t *testing.T) {
	got, err := Decode([]byte(`{"sourceSecret":"demo2-tls","replicas":3,"targets":["demo2"]}`))
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if got.SourceSecret != "demo2-tls" || got.Replicas != 3 || len(got.Targets) != 1 {
		t.Errorf("got %+v, want sourceSecret demo2-tls, replicas 3, one target", got)
	}
}

func TestDecodeIgnoresUnknownFields(t *testing.T) {
	got, err := Decode([]byte(`{"sourceSecret":"demo2-tls","nonsense":true}`))
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if got.SourceSecret != "demo2-tls" {
		t.Errorf("got %+v, want the known field to survive", got)
	}
}
