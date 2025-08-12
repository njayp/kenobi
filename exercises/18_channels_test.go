package exercises

import (
	"testing"
	"time"
)

// Exercise: Channels
// Task: Work with channels for communication between goroutines

// BasicChannel: Create and use a basic channel
func BasicChannel() string {
	// TODO: Create a channel of strings
	// TODO: Send "hello" to the channel in a goroutine
	// TODO: Receive and return the value from the channel
	return ""
}

// BufferedChannel: Work with buffered channels
func BufferedChannel() []int {
	// TODO: Create buffered channel with capacity 3
	// TODO: Send values 1, 2, 3 to channel (without goroutine since it's buffered)
	// TODO: Receive all values and return as slice
	return nil
}

// ChannelRange: Use range to receive from channel
func ChannelRange() []int {
	// TODO: Create channel
	// TODO: In goroutine, send numbers 1-5 to channel, then close it
	// TODO: Use range to receive all values and return as slice
	return nil
}

// ChannelSelect: Use select statement with channels
func ChannelSelect() string {
	// TODO: Create two channels (c1, c2)
	// TODO: In separate goroutines, send "from c1" to c1 after 1ms, "from c2" to c2 after 2ms
	// TODO: Use select to receive from whichever channel is ready first
	// TODO: Return the received message
	return ""
}

// ChannelDirection: Use directional channels
func ChannelDirection() int {
	// TODO: Create channel
	// TODO: Call helper functions with directional channels
	// TODO: Return received value
	return 0
}

// sendOnly: Helper function that only sends to channel
func sendOnly(ch chan<- int) {
	// TODO: Send value 42 to channel
}

// receiveOnly: Helper function that only receives from channel
func receiveOnly(ch <-chan int) int {
	// TODO: Receive and return value from channel
	return 0
}

// ChannelTimeout: Use select with timeout
func ChannelTimeout() string {
	// TODO: Create channel
	// TODO: Use select with:
	//   - case to receive from channel
	//   - case with time.After(50ms) for timeout
	// TODO: Since no one sends to channel, should timeout
	// TODO: Return "timeout" if timeout, otherwise return received value
	return ""
}

func TestBasicChannel(t *testing.T) {
	result := BasicChannel()
	expected := "hello"
	
	if result != expected {
		t.Errorf("BasicChannel() = %q, want %q", result, expected)
	}
}

func TestBufferedChannel(t *testing.T) {
	result := BufferedChannel()
	expected := []int{1, 2, 3}
	
	if len(result) != len(expected) {
		t.Errorf("Length = %d, want %d", len(result), len(expected))
		return
	}
	
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Index %d: got %d, want %d", i, result[i], v)
		}
	}
}

func TestChannelRange(t *testing.T) {
	result := ChannelRange()
	expected := []int{1, 2, 3, 4, 5}
	
	if len(result) != len(expected) {
		t.Errorf("Length = %d, want %d", len(result), len(expected))
		return
	}
	
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Index %d: got %d, want %d", i, result[i], v)
		}
	}
}

func TestChannelSelect(t *testing.T) {
	result := ChannelSelect()
	
	// Should receive from c1 since it sends after 1ms (faster than c2's 2ms)
	expected := "from c1"
	
	if result != expected {
		t.Errorf("ChannelSelect() = %q, want %q", result, expected)
	}
}

func TestChannelDirection(t *testing.T) {
	result := ChannelDirection()
	expected := 42
	
	if result != expected {
		t.Errorf("ChannelDirection() = %d, want %d", result, expected)
	}
}

func TestChannelTimeout(t *testing.T) {
	start := time.Now()
	result := ChannelTimeout()
	duration := time.Since(start)
	
	expected := "timeout"
	if result != expected {
		t.Errorf("ChannelTimeout() = %q, want %q", result, expected)
	}
	
	// Should take approximately 50ms due to timeout
	if duration < 40*time.Millisecond || duration > 100*time.Millisecond {
		t.Errorf("Duration = %v, expected around 50ms", duration)
	}
}

// Non-blocking channel operations
func TestNonBlockingChannel(t *testing.T) {
	// TODO: Students should implement non-blocking channel operations
	// This is a placeholder for teaching select with default case
}

func NonBlockingReceive() (int, bool) {
	// TODO: Create channel, use select with default to try non-blocking receive
	// TODO: Return value and whether receive was successful
	return 0, false
}

func TestNonBlockingReceive(t *testing.T) {
	value, ok := NonBlockingReceive()
	
	// Should not receive anything since no one is sending
	if ok {
		t.Errorf("NonBlockingReceive() received %d, expected no value", value)
	}
}
