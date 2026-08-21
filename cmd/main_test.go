package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"
)

// A fixed port, not an ephemeral one - the same tradeoff spa/justfile makes for
// its own dev server: simpler than plumbing the OS-assigned port back out of
// http.Server, at the cost of a collision if something else is already on it.
const testPort = "18099"

func fixedPort(string) string { return testPort }

func TestRunServesHealthzThenShutsDownCleanly(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var stdout bytes.Buffer

	errCh := make(chan error, 1)
	go func() { errCh <- run(ctx, fixedPort, &stdout) }()

	url := "http://127.0.0.1:" + testPort + "/healthz"
	deadline := time.Now().Add(2 * time.Second)
	var lastErr error
	for time.Now().Before(deadline) {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				t.Fatalf("GET %s = %d, want 200", url, resp.StatusCode)
			}
			lastErr = nil
			break
		}
		lastErr = err
		time.Sleep(10 * time.Millisecond)
	}
	if lastErr != nil {
		t.Fatalf("server never came up: %v", lastErr)
	}

	cancel()

	select {
	case err := <-errCh:
		if err != nil {
			t.Errorf("run() = %v, want nil after a clean shutdown", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("run() did not return within 2s of ctx being canceled")
	}
}

func TestRunLogsStartupAsJSON(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	var stdout bytes.Buffer

	if err := run(ctx, fixedPort, &stdout); err != nil {
		t.Fatalf("run() = %v, want nil", err)
	}

	lines := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(lines) == 0 || lines[0] == "" {
		t.Fatal("stdout is empty, want at least one log line")
	}
	var fields map[string]any
	if err := json.Unmarshal([]byte(lines[0]), &fields); err != nil {
		t.Fatalf("first line is not valid JSON: %v\n%s", err, lines[0])
	}
	if fields["port"] != testPort {
		t.Errorf("fields[port] = %v, want %q logged on startup", fields["port"], testPort)
	}
}
