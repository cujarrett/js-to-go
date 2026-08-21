package cleanup

import (
	"reflect"
	"strings"
	"testing"
)

func TestWithCleanupRunsAfterNormalReturn(t *testing.T) {
	log := &Log{}
	WithCleanup(log, func() { log.add("work") })

	want := []string{"work", "cleanup"}
	if !reflect.DeepEqual(log.lines, want) {
		t.Errorf("log = %v, want %v", log.lines, want)
	}
}

func TestWithCleanupRunsEvenOnPanic(t *testing.T) {
	log := &Log{}
	func() {
		defer func() { recover() }() // keep the test itself alive
		WithCleanup(log, func() {
			log.add("work")
			panic("boom")
		})
	}()

	want := []string{"work", "cleanup"}
	if !reflect.DeepEqual(log.lines, want) {
		t.Errorf("log = %v, want %v - cleanup must run even when work panics", log.lines, want)
	}
}

func TestSafeReturnsNilWhenNothingPanics(t *testing.T) {
	err := Safe(func() {})
	if err != nil {
		t.Errorf("Safe() = %v, want nil", err)
	}
}

func TestSafeTurnsPanicIntoError(t *testing.T) {
	err := Safe(func() { panic("bad object") })
	if err == nil {
		t.Fatal("Safe() = nil, want an error - a panic must not escape")
	}
	if !strings.Contains(err.Error(), "bad object") {
		t.Errorf("err = %q, want it to mention what panicked", err)
	}
}
