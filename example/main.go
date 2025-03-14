package main

import (
	"github.com/ABA-Developer/go-logger"
)

func main() {
	SyncImplementation()
	AsyncImplementation()
}

func SyncImplementation() {
	logger := logger.NewSync("TEST", true)
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
	logger.Info("")
}

func AsyncImplementation() {
	logger := logger.NewAsync("TEST", 10, true)
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
	logger.Flush()
}
