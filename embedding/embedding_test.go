package embedding

import "testing"

func TestHasLabel(t *testing.T) {
	m := Metadata{Name: "web-1", Labels: map[string]string{"env": "prod"}}
	if !m.HasLabel("env", "prod") {
		t.Error("HasLabel(env, prod) = false, want true")
	}
	if m.HasLabel("env", "staging") {
		t.Error("HasLabel(env, staging) = true, want false")
	}
	if m.HasLabel("missing", "x") {
		t.Error("HasLabel(missing, x) = true, want false")
	}
}

func TestServerPromotesMetadataFieldsAndMethods(t *testing.T) {
	s := Server{
		Metadata: Metadata{Name: "web-1", Labels: map[string]string{"env": "prod"}},
		Active:   true,
	}

	// No .Metadata. in front of either - both are promoted onto Server directly.
	if s.Name != "web-1" {
		t.Errorf("s.Name = %q, want web-1", s.Name)
	}
	if !s.HasLabel("env", "prod") {
		t.Error("s.HasLabel(env, prod) = false, want true - HasLabel should be promoted")
	}
}

func TestReconcilerPromotesClientMethod(t *testing.T) {
	r := Reconciler{
		Client: memClient{"web-1": "10.0.0.1"},
		Name:   "server-controller",
	}

	// Get is Client's method, called directly on Reconciler with no r.Client. -
	// exactly how a real Reconciler calls r.Get(ctx, key, obj) via the embedded
	// client.Client, never r.Client.Get(...).
	got, ok := r.Get("web-1")
	if !ok || got != "10.0.0.1" {
		t.Errorf("r.Get(web-1) = %q, %v, want 10.0.0.1, true", got, ok)
	}
}
