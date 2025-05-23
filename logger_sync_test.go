package logger

import (
	"strings"
	"testing"
)

func TestLoggerSync_Info_Color(t *testing.T) {
	logger := NewSync("TEST", false)
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

func TestLoggerSync_Infof_Color(t *testing.T) {
	logger := NewSync("TEST", false)
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

func TestLoggerSync_Warn_Color(t *testing.T) {
	logger := NewSync("TEST", false)
	logger.SetWarnStyle(StyleFgYellow)

	output := CaptureLogOutput(func() {
		logger.Warn("Colored warning")
	})

	if !strings.Contains(output, "\033[33m") {
		t.Errorf("Expected yellow ANSI escape sequence, but got: %q", output)
	}
}

func TestLoggerSync_Warnf_Color(t *testing.T) {
	logger := NewSync("TEST", false)
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

func TestLoggerSync_Error_Color(t *testing.T) {
	logger := NewSync("TEST", false)
	logger.SetErrorStyle(StyleFgRed)

	output := CaptureLogOutput(func() {
		logger.Error("Colored error")
	})

	if !strings.Contains(output, "\033[31m") {
		t.Errorf("Expected red ANSI escape sequence, but got: %q", output)
	}
}

func TestLoggerSync_Errorf_Color(t *testing.T) {
	logger := NewSync("TEST", false)
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

func TestLoggerSync_Debug_Color(t *testing.T) {
	logger := NewSync("TEST", true) // Debug mode enabled
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

func TestLoggerSync_Debugf_Color(t *testing.T) {
	logger := NewSync("TEST", true) // Debug mode enabled
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
