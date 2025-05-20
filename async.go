package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type LoggerAsync struct {
	ch              chan string
	chRaw           chan string
	tag             string
	enDebug         bool
	infoStyle       []int8
	warnStyle       []int8
	errorStyle      []int8
	debugStyle      []int8
	panicStyle      []int8
	fatalStyle      []int8
	writeFileEnable bool
	gateName        string
	file            *os.File
	fileName        string
	path            string
}

// New creates a new Logger instance
// tag: a string that will be displayed in the log message, max 7 characters
// bufferSize: the size of the buffered channel
// debugMode: if true, debug messages will be displayed
// returns a pointer to the Logger instance
// Example:
// logger := logger.New("GPIO", 100, true)
// logger.SetDefaultStyle()
// logger.Info("GPIO handler started")
// log format: [INFO] [TIME] [GPIO]: GPIO handler started
func NewAsync(tag string, bufferSize int, debugMode bool, gateName string) *LoggerAsync {
	lenTag := len(tag)
	if lenTag < 7 {
		tag += strings.Repeat(" ", 7-lenTag)
	}
	tag = "[" + tag + "]"

	// Initial file object
	pathInit := newFolderPath("log_files")
	fileNameInit := fileNameGenerator(gateName)
	fileInit := createAndAppendObject(fileNameInit, pathInit)

	logger := &LoggerAsync{
		ch:              make(chan string, bufferSize), // Buffered channel
		chRaw:           make(chan string, bufferSize), // Buffered channel
		tag:             tag,
		enDebug:         debugMode,
		writeFileEnable: false,
		gateName:        gateName,
		fileName:        fileNameInit,
		file:            fileInit,
		path:            pathInit,
	}
	logger.init()
	return logger
}

// Async logging function that runs in a separate goroutine
func (l *LoggerAsync) init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0) // Disable the default timestamp and log prefix
	go func() {
		for logMsg := range l.ch {
			log.Print(logMsg)
		}
	}()

	go func() {
		for logMsgRaw := range l.chRaw {
			l.writeLog(logMsgRaw)
		}
	}()
}

func (l *LoggerAsync) ChangeFileRoutine(hour int, minute int) {
	HOUR := hour
	MINUTE := minute
	go func() {
		for range time.Tick(1 * time.Minute) {
			hours, minutes, _ := time.Now().Clock()
			if hours == HOUR && minutes == MINUTE {
				// Close first previous object file
				l.file.Close()

				// Create new file object with the append mode
				l.fileName = fileNameGenerator(l.gateName)
				l.file = createAndAppendObject(l.fileName, l.path)
			}
		}
	}()

}

func (l *LoggerAsync) writeLog(msg string) {
	if l.writeFileEnable {
		l.file.WriteString(msg + "\n")
	}
}

func (l *LoggerAsync) SetWriteFilesEnable(enable bool) {
	l.writeFileEnable = enable
}

func (l *LoggerAsync) SetPath(path string) {
	l.file.Close()

	// New path with folder creation
	l.path = path
	newFolderPath(path)

	// Initial file object
	l.fileName = fileNameGenerator(l.gateName)
	l.file = createAndAppendObject(l.fileName, l.path)
}

func (l *LoggerAsync) applyStyle(str string, styles ...int8) string {
	resStr := "\033["

	for i := range len(styles) {
		if i == 0 {
			resStr += fmt.Sprintf("%d", styles[i])
		} else {
			resStr += fmt.Sprintf(";%d", styles[i])
		}
	}
	resStr += fmt.Sprintf("m%s\033[0m", str)

	return resStr
}

func (l *LoggerAsync) getTime() string {
	return "[" + time.Now().Format("2006-01-02 15:04:05.000") + "] "
}

func (l *LoggerAsync) Flush() {
	close(l.ch)
	close(l.chRaw)
	for logMsg := range l.ch {
		log.Print(logMsg)
	}
	for logMsgRaw := range l.chRaw {
		log.Print(logMsgRaw)
	}
}

func (l *LoggerAsync) SetInfoStyle(styles ...int8) {
	for i := range len(styles) {
		l.infoStyle = append(l.infoStyle, styles[i])
	}
}

func (l *LoggerAsync) SetWarnStyle(styles ...int8) {
	for i := range len(styles) {
		l.warnStyle = append(l.warnStyle, styles[i])
	}
}

func (l *LoggerAsync) SetErrorStyle(styles ...int8) {
	for i := range len(styles) {
		l.errorStyle = append(l.errorStyle, styles[i])
	}
}

func (l *LoggerAsync) SetDebugStyle(styles ...int8) {
	for i := range len(styles) {
		l.debugStyle = append(l.debugStyle, styles[i])
	}
}

func (l *LoggerAsync) SetPanicStyle(styles ...int8) {
	for i := range len(styles) {
		l.panicStyle = append(l.panicStyle, styles[i])
	}
}

func (l *LoggerAsync) SetFatalStyle(styles ...int8) {
	for i := range len(styles) {
		l.fatalStyle = append(l.fatalStyle, styles[i])
	}
}

func (l *LoggerAsync) SetDefaultStyle() {
	l.SetInfoStyle(StyleFgWhite)
	l.SetWarnStyle(StyleFgYellow)
	l.SetErrorStyle(StyleFgRed)
	l.SetDebugStyle(StyleFontItalic, StyleFontComment)
	l.SetPanicStyle(StyleFontBold, StyleFgBlack, StyleBgMagenta)
	l.SetFatalStyle(StyleFontBold, StyleFgBlack, StyleBgRed)
}

// LOG FORMAT
// [INFO] [TIME] [TAG]: message

func (l *LoggerAsync) Info(a ...any) {
	msg := l.getTime() + infoKey + l.tag + ": " + fmt.Sprint(a...)
	l.chRaw <- msg
	msg = l.applyStyle(msg, l.infoStyle...)
	l.ch <- msg
}

func (l *LoggerAsync) Infof(format string, a ...any) {
	msg := l.getTime() + infoKey + l.tag + ": " + fmt.Sprintf(format, a...)
	l.chRaw <- msg
	msg = l.applyStyle(msg, l.infoStyle...)
	l.ch <- msg
}

func (l *LoggerAsync) Warn(a ...any) {
	msg := l.getTime() + warnKey + l.tag + ": " + fmt.Sprint(a...)
	l.chRaw <- msg
	msg = l.applyStyle(msg, l.warnStyle...)
	l.ch <- msg
}

func (l *LoggerAsync) Warnf(format string, a ...any) {
	msg := l.getTime() + warnKey + l.tag + ": " + fmt.Sprintf(format, a...)
	l.chRaw <- msg
	msg = l.applyStyle(msg, l.warnStyle...)
	l.ch <- msg
}

func (l *LoggerAsync) Error(a ...any) {
	msg := l.getTime() + errorKey + l.tag + ": " + fmt.Sprint(a...)
	l.chRaw <- msg
	msg = l.applyStyle(msg, l.errorStyle...)
	l.ch <- msg
}

func (l *LoggerAsync) Errorf(format string, a ...any) {
	msg := l.getTime() + errorKey + l.tag + ": " + fmt.Sprintf(format, a...)
	l.chRaw <- msg
	msg = l.applyStyle(msg, l.errorStyle...)
	l.ch <- msg
}

func (l *LoggerAsync) Debug(a ...any) {
	if l.enDebug {
		msg := l.getTime() + debugKey + l.tag + ": " + fmt.Sprint(a...)
		l.chRaw <- msg
		msg = l.applyStyle(msg, l.debugStyle...)
		l.ch <- msg
	}
}

func (l *LoggerAsync) Debugf(format string, a ...any) {
	if l.enDebug {
		msg := l.getTime() + debugKey + l.tag + ": " + fmt.Sprintf(format, a...)
		l.chRaw <- msg
		msg = l.applyStyle(msg, l.debugStyle...)
		l.ch <- msg
	}
}

func (l *LoggerAsync) Panic(a ...any) {
	msg := l.getTime() + panicKey + l.tag + ": " + fmt.Sprint(a...)
	l.chRaw <- msg
	msg = l.applyStyle(msg, l.panicStyle...)
	l.ch <- msg
}

func (l *LoggerAsync) Panicf(format string, a ...any) {
	msg := l.getTime() + panicKey + l.tag + ": " + fmt.Sprintf(format, a...)
	l.chRaw <- msg
	msg = l.applyStyle(msg, l.panicStyle...)
	l.ch <- msg
}

func (l *LoggerAsync) Fatal(a ...any) {
	msg := l.getTime() + fatalKey + l.tag + ": " + fmt.Sprint(a...)
	l.chRaw <- msg
	msg = l.applyStyle(msg, l.fatalStyle...)
	l.ch <- msg
}

func (l *LoggerAsync) Fatalf(format string, a ...any) {
	msg := l.getTime() + fatalKey + l.tag + ": " + fmt.Sprintf(format, a...)
	l.chRaw <- msg
	msg = l.applyStyle(msg, l.fatalStyle...)
	l.ch <- msg
}
