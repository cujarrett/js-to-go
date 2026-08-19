// Package service is a small HTTP API built the way every Go service in this
// workspace is built: standard library only, one dependency behind an interface.
package service

import (
	"net/http"
)

// Store is the one dependency. Taking an interface here is what makes the
// handlers testable without a database.
type Store interface {
	Get(name string) (Server, bool)
	List() []Server
}

// Server is one machine.
type Server struct {
	Name   string `json:"name"`
	Slot   string `json:"slot"`
	Active bool   `json:"active"`
}

// API holds the handlers' dependencies.
type API struct {
	store Store
}

// New returns an API backed by store.
func New(store Store) *API {
	return &API{store: store}
}

// Routes returns the API's handler. Go 1.22 and later allow the method and path
// parameters in the pattern itself, so no router library is needed.
func (a *API) Routes() http.Handler {
	mux := http.NewServeMux()
	// TODO: register
	//   GET /healthz          -> a.handleHealth
	//   GET /servers          -> a.handleList
	//   GET /servers/{name}   -> a.handleGet
	return mux
}

// handleHealth reports that the process is up. Every service here has one.
func (a *API) handleHealth(w http.ResponseWriter, r *http.Request) {
	// TODO: 200 with the body "ok"
}

// handleList writes every server as a JSON array.
func (a *API) handleList(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// handleGet writes one server as JSON, or 404 when it does not exist.
// The name comes from the path: r.PathValue("name").
func (a *API) handleGet(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// writeJSON is the one place that sets the content type and encodes. Keeping it
// in a helper stops three handlers disagreeing about headers.
func writeJSON(w http.ResponseWriter, status int, v any) {
	// TODO
}
