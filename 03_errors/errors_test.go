package errors

import (
	stderrors "errors"
	"strings"
	"testing"
)

func store() *Store {
	return NewStore(map[string]string{"demo1": "web-1"})
}

func TestGetFound(t *testing.T) {
	got, err := store().Get("demo1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got != "web-1" {
		t.Errorf("Get() = %q, want web-1", got)
	}
}

func TestGetMissingReturnsSentinel(t *testing.T) {
	_, err := store().Get("demo9")
	if !stderrors.Is(err, ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestDescribeWrapsWithContext(t *testing.T) {
	_, err := store().Describe("demo9")
	if err == nil {
		t.Fatal("Describe() = nil, want an error")
	}
	if !stderrors.Is(err, ErrNotFound) {
		t.Errorf("err = %v, want it to still unwrap to ErrNotFound - use %%w, not %%v", err)
	}
	if !strings.Contains(err.Error(), "demo9") {
		t.Errorf("err = %q, want the key in the message", err)
	}
}

func TestRequireCarriesTheKey(t *testing.T) {
	err := store().Require("demo9")

	var missing *MissingKeyError
	if !stderrors.As(err, &missing) {
		t.Fatalf("err = %v, want a *MissingKeyError", err)
	}
	if missing.Key != "demo9" {
		t.Errorf("Key = %q, want demo9", missing.Key)
	}
}

func TestRequireFoundIsNil(t *testing.T) {
	if err := store().Require("demo1"); err != nil {
		t.Errorf("Require() = %v, want nil", err)
	}
}
