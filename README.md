# Logger for Go

A simple, customizable logger for Go with both synchronous and asynchronous logging modes, colored output, and various log levels.

## Features
- Supports **sync** and **async** logging
- ANSI color formatting for log levels
- Supports **Info, Debug, Warn, Error, Panic, Fatal** messages
- Buffered channel for async logging with **Flush()** support
- Configurable log styles
- Tagged log messages

## Installation

```sh
go get github.com/ABA-Developer/go-logger
```

## Usage

### Synchronous Logger

```go
package main

import (
	"github.com/ABA-Developer/go-logger"
)

func main() {
	SyncImplementation()
}

func SyncImplementation() {
	logger := logger.NewSync("TEST", true) // Tag: "TEST", Debug mode enabled
	logger.SetDefaultStyle()
	logger.Debug("Sync logger started")

	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			logger.Error(i)
		} else if i%5 == 0 {
			logger.Warn(i)
		} else {
			logger.Info(i)
		}
	}
}
```

### Asynchronous Logger

```go
package main

import (
	"github.com/ABA-Developer/go-logger"
)

func main() {
	AsyncImplementation()
}

func AsyncImplementation() {
	logger := logger.NewAsync("TEST", 10, true) // Tag: "TEST", Buffer size: 10, Debug mode enabled
	logger.SetDefaultStyle()
	logger.Debug("Async logger started")

	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			logger.Error(i)
		} else if i%5 == 0 {
			logger.Warn(i)
		} else {
			logger.Info(i)
		}
	}

	logger.Flush() // Ensure all logs are printed before exit
}
```

## API Reference

### Logger Creation

- `NewSync(tag string, debugMode bool) *LoggerSync`
- `NewAsync(tag string, bufferSize int, debugMode bool) *LoggerAsync`

### Logging Methods

- `logger.Info(a ...any)`
- `logger.Infof(format string, a ...any)`
- `logger.Debug(a ...any)`
- `logger.Debugf(format string, a ...any)`
- `logger.Warn(a ...any)`
- `logger.Warnf(format string, a ...any)`
- `logger.Error(a ...any)`
- `logger.Errorf(format string, a ...any)`
- `logger.Panic(a ...any)`
- `logger.Panicf(format string, a ...any)`
- `logger.Fatal(a ...any)`
- `logger.Fatalf(format string, a ...any)`

### Additional Methods

- `logger.Flush()` (only for async logger)
- `logger.SetInfoStyle(styles ...int8)`
- `logger.SetWarnStyle(styles ...int8)`
- `logger.SetErrorStyle(styles ...int8)`
- `logger.SetDebugStyle(styles ...int8)`
- `logger.SetPanicStyle(styles ...int8)`
- `logger.SetFatalStyle(styles ...int8)`
- `logger.SetDefaultStyle()`

## Log Output Format

```
[INFO ] [YYYY-MM-DD HH:MM:SS.mmm] [TAG]: message
[WARN ] [YYYY-MM-DD HH:MM:SS.mmm] [TAG]: message
[ERROR] [YYYY-MM-DD HH:MM:SS.mmm] [TAG]: message
[DEBUG] [YYYY-MM-DD HH:MM:SS.mmm] [TAG]: message
```

## License
MIT

