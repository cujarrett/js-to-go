package logging

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestLogRequestIsOneJSONLine(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)

	LogRequest(logger, "GET", "/servers", 200)

	line := strings.TrimSpace(buf.String())
	if strings.Count(line, "\n") != 0 {
		t.Fatalf("output = %q, want exactly one line", buf.String())
	}

	var fields map[string]any
	if err := json.Unmarshal([]byte(line), &fields); err != nil {
		t.Fatalf("line is not valid JSON: %v\n%s", err, line)
	}

	if fields["method"] != "GET" || fields["path"] != "/servers" {
		t.Errorf("fields = %v, want method GET and path /servers", fields)
	}
	if status, ok := fields["status"].(float64); !ok || status != 200 {
		t.Errorf("fields[status] = %v, want the number 200, not a formatted string", fields["status"])
	}
}

func TestWithRequestIDStampsEveryLine(t *testing.T) {
	var buf bytes.Buffer
	base := New(&buf)
	scoped := WithRequestID(base, "req-42")

	LogRequest(scoped, "GET", "/servers", 200)
	LogRequest(scoped, "GET", "/servers/web-1", 200)

	for i, line := range strings.Split(strings.TrimSpace(buf.String()), "\n") {
		var fields map[string]any
		if err := json.Unmarshal([]byte(line), &fields); err != nil {
			t.Fatalf("line %d is not valid JSON: %v", i, err)
		}
		if fields["request_id"] != "req-42" {
			t.Errorf("line %d: request_id = %v, want req-42", i, fields["request_id"])
		}
	}
}

func TestBaseLoggerUnaffected(t *testing.T) {
	var buf bytes.Buffer
	base := New(&buf)
	_ = WithRequestID(base, "req-42") // scoped copy, base must not change

	LogRequest(base, "GET", "/healthz", 200)

	var fields map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(buf.Bytes()), &fields); err != nil {
		t.Fatalf("not valid JSON: %v", err)
	}
	if _, has := fields["request_id"]; has {
		t.Errorf("fields = %v, want the base logger to carry no request_id", fields)
	}
}
