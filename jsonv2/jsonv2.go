// Package jsonv2 is one difference between encoding/json and its v2 successor,
// isolated. On Go 1.26, v2 only builds under GOEXPERIMENT=jsonv2, which
// `just test jsonv2` and `just ci` set for you.
package jsonv2

import "time"

// Record is a resource with a sync timestamp - the same shape a controller's
// status subresource carries.
type Record struct {
	Name string `json:"name"`
	// v1's omitempty never treats a struct as empty, zero value or not - only nil
	// pointers, empty maps/slices and zero scalars qualify. A zero time.Time comes
	// back "0001-01-01T00:00:00Z" no matter what you write here.
	// TODO: use v2's "omitzero" instead, which asks the type itself whether it is
	// zero. time.Time already implements IsZero, so this just works.
	LastSynced time.Time `json:"lastSynced"`
}

// Encode returns r as JSON, using encoding/json/v2.
func Encode(r Record) ([]byte, error) {
	// TODO
	return nil, nil
}
