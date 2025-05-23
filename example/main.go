package main

import (
	"time"

	"github.com/ABA-Developer/go-logger"
)

func main() {
	SyncImplementation()
	AsyncImplementation()
}

func SyncImplementation() {
	logger := logger.NewSync("TEST", true)
	logger.SetWriteFilesEnable("log_files", "ABA11")
	logger.ChangeFileRoutine(00, 00)
	logger.SetDefaultStyle()
	logger.Debug("Sync logger started")
	logMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Logging test for long message.12345"
	i := 0
	for {
		if i%3 == 0 {
			logger.Error(i, logMessage)
		} else if i%5 == 0 {
			logger.Warn(i, logMessage)
		} else {
			logger.Info(i, logMessage)
		}
		i++
		time.Sleep(100 * time.Millisecond)
		if i > 200 {
			break
		}
	}
	logger.Info("")
}

func AsyncImplementation() {
	logger := logger.NewAsync("TEST", 10, true)
	logger.SetWriteFilesEnable("log_files", "ABA11")
	logger.ChangeFileRoutine(00, 00)
	logger.SetDefaultStyle()
	logger.Debug("Async logger started")
	logMessage := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Logging test for long message.12345"
	i := 0
	for {
		if i%3 == 0 {
			logger.Error(i, logMessage)
		} else if i%5 == 0 {
			logger.Warn(i, logMessage)
		} else {
			logger.Info(i, logMessage)
		}
		i++
		time.Sleep(100 * time.Millisecond)
		if i > 200 {
			break
		}
	}
	logger.Flush()
}
