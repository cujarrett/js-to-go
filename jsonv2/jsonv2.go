// Package jsonv2 is one difference between encoding/json and its v2 successor,
// isolated. v2 is real but experimental: this module only builds and tests under
// GOEXPERIMENT=jsonv2, which `just test 06` and `just ci` set for you.
package jsonv2

import "time"

// Record is a resource with a sync timestamp, the shape a controller's status
// subresource carries.
type Record struct {
	Name string `json:"name"`
	// v1's omitempty never treats a struct as empty, whatever its value.
	// Only nil pointers, empty maps/slices and zero scalars qualify, so a zero
	// time.Time still encodes as "0001-01-01T00:00:00Z".
	// TODO: use v2's "omitzero", which asks the type whether it is zero.
	// time.Time already implements IsZero, so the tag is the only change needed.
	LastSynced time.Time `json:"lastSynced"`
}

// Encode returns r as JSON, using encoding/json/v2.
func Encode(r Record) ([]byte, error) {
	// TODO
	return nil, nil
}
