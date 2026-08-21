package taskrunner

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestSlowOpFinishes(t *testing.T) {
	if err := SlowOp(context.Background(), time.Millisecond); err != nil {
		t.Errorf("SlowOp() = %v, want nil", err)
	}
}

func TestSlowOpTimesOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	err := SlowOp(ctx, 200*time.Millisecond)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err = %v, want context.DeadlineExceeded", err)
	}
}

func TestSlowOpCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := SlowOp(ctx, 200*time.Millisecond)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
}

func TestRequestIDRoundTrips(t *testing.T) {
	ctx := WithRequestID(context.Background(), "req-42")
	got, ok := RequestID(ctx)
	if !ok || got != "req-42" {
		t.Errorf("RequestID() = %q, %v, want req-42, true", got, ok)
	}
}

func TestRequestIDAbsentByDefault(t *testing.T) {
	_, ok := RequestID(context.Background())
	if ok {
		t.Error("RequestID() ok = true on a bare context, want false")
	}
}
