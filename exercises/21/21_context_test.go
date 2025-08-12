package exercises

import (
	"context"
	"testing"
	"time"
)

// Exercise: Context
// Task: Use context for cancellation, timeouts, and value passing

// WithTimeout: Use context with timeout
func WithTimeout(duration time.Duration) (string, error) {
	// TODO: Create context with timeout
	// TODO: Use select to wait for either:
	//   - A channel that sends message after duration*2 (should timeout)
	//   - Context cancellation
	// TODO: Return "completed" if no timeout, error if timeout
	return "", nil
}

// WithCancel: Use context with manual cancellation
func WithCancel() string {
	// TODO: Create cancellable context
	// TODO: Launch goroutine that waits for context cancellation
	// TODO: Cancel context after 10ms
	// TODO: Return "cancelled" when context is cancelled
	return ""
}

// WithValue: Pass values through context
func WithValue(key, value string) string {
	// TODO: Create context with value
	// TODO: Call helper function that retrieves value from context
	// TODO: Return the retrieved value
	return ""
}

// getValue: Helper function to retrieve value from context
func getValue(ctx context.Context, key string) string {
	// TODO: Get value from context, return empty string if not found
	return ""
}

// ChainedContext: Demonstrate context chaining
func ChainedContext() (string, string) {
	// TODO: Create parent context with value "user"="alice"
	// TODO: Create child context with timeout of 100ms
	// TODO: Create grandchild context with value "role"="admin"
	// TODO: Return both values retrieved from grandchild context
	return "", ""
}

// HTTPLikeOperation: Simulate HTTP operation with context
func HTTPLikeOperation(ctx context.Context) error {
	// TODO: Simulate long-running operation
	// TODO: Check context cancellation during operation
	// TODO: Return context.Canceled if cancelled, nil if completed
	select {
	case <-time.After(50 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func TestWithTimeout(t *testing.T) {
	// Test timeout (should timeout since we wait for duration*2)
	start := time.Now()
	result, err := WithTimeout(25 * time.Millisecond)
	duration := time.Since(start)
	
	if err == nil {
		t.Error("WithTimeout should return timeout error")
	}
	if result != "" {
		t.Errorf("WithTimeout result = %q, want empty string on timeout", result)
	}
	
	// Should timeout around 25ms
	if duration < 20*time.Millisecond || duration > 50*time.Millisecond {
		t.Errorf("Timeout duration = %v, expected around 25ms", duration)
	}
}

func TestWithCancel(t *testing.T) {
	start := time.Now()
	result := WithCancel()
	duration := time.Since(start)
	
	expected := "cancelled"
	if result != expected {
		t.Errorf("WithCancel() = %q, want %q", result, expected)
	}
	
	// Should complete in around 10ms
	if duration < 5*time.Millisecond || duration > 25*time.Millisecond {
		t.Errorf("Cancel duration = %v, expected around 10ms", duration)
	}
}

func TestWithValue(t *testing.T) {
	result := WithValue("test-key", "test-value")
	expected := "test-value"
	
	if result != expected {
		t.Errorf("WithValue() = %q, want %q", result, expected)
	}
}

func TestChainedContext(t *testing.T) {
	user, role := ChainedContext()
	
	if user != "alice" {
		t.Errorf("User = %q, want %q", user, "alice")
	}
	if role != "admin" {
		t.Errorf("Role = %q, want %q", role, "admin")
	}
}

func TestHTTPLikeOperation(t *testing.T) {
	// Test successful completion
	ctx := context.Background()
	err := HTTPLikeOperation(ctx)
	if err != nil {
		t.Errorf("HTTPLikeOperation with background context returned error: %v", err)
	}
	
	// Test cancellation
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()
	
	err = HTTPLikeOperation(ctx)
	if err != context.Canceled {
		t.Errorf("HTTPLikeOperation with cancelled context returned %v, want %v", err, context.Canceled)
	}
}

// Advanced: Context with deadline
func TestContextDeadline(t *testing.T) {
	// TODO: Students should implement context with deadline
}

func WithDeadline() error {
	// TODO: Create context with deadline 30ms from now
	// TODO: Try to do work that takes 50ms
	// TODO: Return context error if deadline exceeded
	return nil
}
