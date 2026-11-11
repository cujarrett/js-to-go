// Package jsonv2 is one difference between encoding/json and its v2 successor,
// isolated. v2 is stable from Go 1.27, so go.mod asks for 1.27.
package jsonv2

import (
	"encoding/json/v2"
	"time"
)

// Record is a resource with a sync timestamp - the same shape a controller's
// status subresource carries.
type Record struct {
	Name string `json:"name"`
	// v1's omitempty never treats a struct as empty, zero value or not - only nil
	// pointers, empty maps/slices and zero scalars qualify. A zero time.Time comes
	// back "0001-01-01T00:00:00Z" no matter what you write here.
	// omitzero asks the type itself. time.Time implements IsZero, so a zero time is dropped.
	LastSynced time.Time `json:"lastSynced,omitzero"`
}

// Encode returns r as JSON, using encoding/json/v2.
func Encode(r Record) ([]byte, error) {
	return json.Marshal(r)
}
