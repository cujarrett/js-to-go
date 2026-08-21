package generics

import (
	"reflect"
	"testing"
)

func TestNamesWorksOnServers(t *testing.T) {
	servers := []Server{{name: "web-1"}, {name: "web-2"}}
	got := Names(servers)
	want := []string{"web-1", "web-2"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Names(servers) = %v, want %v", got, want)
	}
}

func TestNamesWorksOnPodsTooWithNoChangeToNames(t *testing.T) {
	pods := []Pod{{podName: "coredns-abc"}, {podName: "coredns-def"}}
	got := Names(pods)
	want := []string{"coredns-abc", "coredns-def"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Names(pods) = %v, want %v", got, want)
	}
}

func TestIntSet(t *testing.T) {
	s := NewSet[int]()
	s.Add(7)
	s.Add(7) // adding twice is still just present once

	if !s.Has(7) {
		t.Error("Has(7) = false, want true")
	}
	if s.Has(8) {
		t.Error("Has(8) = true, want false")
	}
}

func TestStringSet(t *testing.T) {
	// Same Set[T], a different T - nothing in Set changed to support this.
	s := NewSet[string]()
	s.Add("demo1")

	if !s.Has("demo1") {
		t.Error(`Has("demo1") = false, want true`)
	}
	if s.Has("demo2") {
		t.Error(`Has("demo2") = true, want false`)
	}
}
