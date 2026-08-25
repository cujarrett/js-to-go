package collections

import (
	"reflect"
	"testing"
)

func inventory() []Server {
	return []Server{
		{Name: "web-1", Slot: "demo1"},
		{Name: "web-2", Slot: "demo3"},
		{Name: "web-3", Slot: "demo1"},
	}
}

func TestNames(t *testing.T) {
	want := []string{"web-1", "web-2", "web-3"}
	if got := Names(inventory()); !reflect.DeepEqual(got, want) {
		t.Errorf("Names() = %v, want %v", got, want)
	}
}

func TestNamesOnEmptyInput(t *testing.T) {
	if got := Names(nil); len(got) != 0 {
		t.Errorf("Names(nil) = %v, want empty", got)
	}
}

func TestSlotSet(t *testing.T) {
	set := SlotSet(inventory())
	if !set["demo1"] || !set["demo3"] {
		t.Errorf("set = %v, want demo1 and demo3 present", set)
	}
	if set["demo5"] {
		t.Error("set reports demo5 present, want absent")
	}
}

func TestCountBySlot(t *testing.T) {
	counts := CountBySlot(inventory())
	if counts["demo1"] != 2 || counts["demo3"] != 1 {
		t.Errorf("counts = %v, want demo1:2 demo3:1", counts)
	}
	if counts["missing"] != 0 {
		t.Errorf("missing key = %d, want the zero value 0", counts["missing"])
	}
}

func TestActivateAll(t *testing.T) {
	servers := inventory()
	ActivateAll(servers)
	for _, s := range servers {
		if !s.Active {
			t.Fatalf("%s is not Active - ranging copies each element", s.Name)
		}
	}
}

func TestUnselected(t *testing.T) {
	keep := map[string]bool{"demo1": true}
	want := []string{"web-2"}
	if got := Unselected(inventory(), keep); !reflect.DeepEqual(got, want) {
		t.Errorf("Unselected() = %v, want %v", got, want)
	}
}

func TestAddServerReturnsTheGrownSlice(t *testing.T) {
	servers := inventory()
	grown := AddServer(servers, Server{Name: "web-4", Slot: "demo1"})

	if len(grown) != len(servers)+1 {
		t.Errorf("len(grown) = %d, want %d", len(grown), len(servers)+1)
	}
	// servers is still length 3 here, whatever AddServer does - append cannot resize
	// the caller's own variable, only the caller reassigning the return value can.
	// Not asserted: true even of a broken AddServer, so it can never be the red test.
}
