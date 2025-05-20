package logger

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"time"
)

// CaptureLogOutput temporarily redirects log output to a buffer
func CaptureLogOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)      // Redirect logs to buffer
	defer log.SetOutput(nil) // Restore default output

	f() // Run the function that triggers log output

	// Wait for async logging to finish
	time.Sleep(1 * time.Millisecond)

	return buf.String()
}

func TestLoggerAsync_Info_Color(t *testing.T) {
	logger := NewAsync("TEST", 10, false, "ABA11")
	logger.SetInfoStyle(StyleFgGreen) // Expect green ANSI color

	output := CaptureLogOutput(func() {
		logger.Info("Colored info message")
	})

	if !strings.Contains(output, "\033[32m") {
		t.Errorf("Expected green ANSI escape sequence, but got: %q", output)
	}

	if !strings.Contains(output, "Colored info message") {
		t.Errorf("Expected log message but got: %q", output)
	}
}

func TestLoggerAsync_Infof_Color(t *testing.T) {
	logger := NewAsync("TEST", 10, false, "ABA11")
	logger.SetInfoStyle(StyleFgGreen)

	output := CaptureLogOutput(func() {
		logger.Infof("Formatted %s message", "info")
	})

	if !strings.Contains(output, "\033[32m") {
		t.Errorf("Expected green ANSI escape sequence, but got: %q", output)
	}

	if !strings.Contains(output, "Formatted info message") {
		t.Errorf("Expected formatted message but got: %q", output)
	}
}

func TestLoggerAsync_Warn_Color(t *testing.T) {
	logger := NewAsync("TEST", 10, false, "ABA11")
	logger.SetWarnStyle(StyleFgYellow)

	output := CaptureLogOutput(func() {
		logger.Warn("Colored warning")
	})

	if !strings.Contains(output, "\033[33m") {
		t.Errorf("Expected yellow ANSI escape sequence, but got: %q", output)
	}
}

func TestLoggerAsync_Warnf_Color(t *testing.T) {
	logger := NewAsync("TEST", 10, false, "ABA11")
	logger.SetWarnStyle(StyleFgYellow)

	output := CaptureLogOutput(func() {
		logger.Warnf("Formatted %s warning", "yellow")
	})

	if !strings.Contains(output, "\033[33m") {
		t.Errorf("Expected yellow ANSI escape sequence, but got: %q", output)
	}

	if !strings.Contains(output, "Formatted yellow warning") {
		t.Errorf("Expected formatted message but got: %q", output)
	}
}

func TestLoggerAsync_Error_Color(t *testing.T) {
	logger := NewAsync("TEST", 10, false, "ABA11")
	logger.SetErrorStyle(StyleFgRed)

	output := CaptureLogOutput(func() {
		logger.Error("Colored error")
	})

	if !strings.Contains(output, "\033[31m") {
		t.Errorf("Expected red ANSI escape sequence, but got: %q", output)
	}
}

func TestLoggerAsync_Errorf_Color(t *testing.T) {
	logger := NewAsync("TEST", 10, false, "ABA11")
	logger.SetErrorStyle(StyleFgRed)

	output := CaptureLogOutput(func() {
		logger.Errorf("Formatted %s error", "red")
	})

	if !strings.Contains(output, "\033[31m") {
		t.Errorf("Expected red ANSI escape sequence, but got: %q", output)
	}

	if !strings.Contains(output, "Formatted red error") {
		t.Errorf("Expected formatted message but got: %q", output)
	}
}

func TestLoggerAsync_Debug_Color(t *testing.T) {
	logger := NewAsync("TEST", 10, true, "ABA11") // Debug mode enabled
	logger.SetDebugStyle(StyleFgCyan)

	output := CaptureLogOutput(func() {
		logger.Debug("Colored debug message")
	})

	if !strings.Contains(output, "\033[36m") {
		t.Errorf("Expected cyan ANSI escape sequence, but got: %q", output)
	}

	if !strings.Contains(output, "Colored debug message") {
		t.Errorf("Expected debug message but got: %q", output)
	}
}

func TestLoggerAsync_Debugf_Color(t *testing.T) {
	logger := NewAsync("TEST", 10, true, "ABA11") // Debug mode enabled
	logger.SetDebugStyle(StyleFgCyan)

	output := CaptureLogOutput(func() {
		logger.Debugf("Formatted %s debug", "cyan")
	})

	if !strings.Contains(output, "\033[36m") {
		t.Errorf("Expected cyan ANSI escape sequence, but got: %q", output)
	}

	if !strings.Contains(output, "Formatted cyan debug") {
		t.Errorf("Expected formatted debug message but got: %q", output)
	}
}

// panic and fatal tests are not included because they will terminate the test process
