package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// fakeStore is the whole reason Store is an interface. No database, no mocking
// library, just a type that satisfies the same methods.
type fakeStore struct {
	servers []Server
}

func (f *fakeStore) Get(name string) (Server, bool) {
	for _, s := range f.servers {
		if s.Name == name {
			return s, true
		}
	}
	return Server{}, false
}

func (f *fakeStore) List() []Server { return f.servers }

func testAPI() http.Handler {
	return New(&fakeStore{servers: []Server{
		{Name: "web-1", Slot: "demo1", Active: true},
		{Name: "web-2", Slot: "demo3"},
	}}).Routes()
}

func do(t *testing.T, method, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	testAPI().ServeHTTP(rec, httptest.NewRequest(method, path, nil))
	return rec
}

func TestHealthz(t *testing.T) {
	rec := do(t, http.MethodGet, "/healthz")
	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", rec.Code)
	}
	if strings.TrimSpace(rec.Body.String()) != "ok" {
		t.Errorf("body = %q, want ok", rec.Body.String())
	}
}

func TestListReturnsJSON(t *testing.T) {
	rec := do(t, http.MethodGet, "/servers")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Errorf("Content-Type = %q, want application/json", ct)
	}

	var got []Server
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode: %v (body %s)", err, rec.Body)
	}
	if len(got) != 2 {
		t.Errorf("got %d servers, want 2", len(got))
	}
}

func TestGetOne(t *testing.T) {
	rec := do(t, http.MethodGet, "/servers/web-1")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}

	var got Server
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.Name != "web-1" || got.Slot != "demo1" {
		t.Errorf("got %+v, want web-1 in demo1", got)
	}
}

func TestGetMissingIs404(t *testing.T) {
	if rec := do(t, http.MethodGet, "/servers/nope"); rec.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", rec.Code)
	}
}

func TestWrongMethodIs405(t *testing.T) {
	if rec := do(t, http.MethodPost, "/servers"); rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want 405 - register the method in the pattern", rec.Code)
	}
}
