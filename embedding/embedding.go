// Package embedding is what every Kubernetes object type and every controller
// struct you will read leans on: metav1.ObjectMeta is embedded in every API
// type, and a Reconciler almost always embeds client.Client.
package embedding

// Metadata mirrors the shape of metav1.ObjectMeta closely enough for this
// exercise - a name and a set of labels every resource in this repo carries.
type Metadata struct {
	Name   string
	Labels map[string]string
}

// HasLabel reports whether key is present with that exact value.
func (m Metadata) HasLabel(key, value string) bool {
	// TODO
	return false
}

// Server is a machine. Note the field with no name - that is the embed. It is
// not "Server has a Metadata field called Metadata"; it is "Server has-a
// Metadata, and Metadata's own methods and fields become Server's too."
type Server struct {
	Metadata
	Active bool
}

// Reconciler is the shape of a real controller - it embeds client.Client, so
// every method Client has (Get, List, Update, ...) becomes a method of
// Reconciler too, with no line written here to forward any of them.
type Reconciler struct {
	Client
	Name string
}

// Client stands in for client.Client - enough of it to prove the point.
type Client interface {
	Get(key string) (string, bool)
}

// memClient is a trivial Client for the test to hand a Reconciler.
type memClient map[string]string

func (m memClient) Get(key string) (string, bool) {
	v, ok := m[key]
	return v, ok
}
