package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type LoggerSync struct {
	tag              string
	enDebug          bool
	infoStyle        []int8
	warnStyle        []int8
	errorStyle       []int8
	debugStyle       []int8
	panicStyle       []int8
	fatalStyle       []int8
	writeFilesEnable bool
	gateName         string
	file             *os.File
	filesName        string
}

// New creates a new Logger instance
// tag: a string that will be displayed in the log message, max 7 characters
// debugMode: if true, debug messages will be displayed
// returns a pointer to the Logger instance
// Example:
// logger := logger.New("GPIO", true)
// logger.SetDefaultStyle()
// logger.Info("GPIO handler started")
// log format: [INFO] [TIME] [GPIO]: GPIO handler started
func NewSync(tag string, debugMode bool, gateName string) *LoggerSync {
	lenTag := len(tag)
	if lenTag < 7 {
		tag += strings.Repeat(" ", 7-lenTag)
	}
	tag = "[" + tag + "]"

	// Initial file object
	fileNameInit := fileNameGenerator(gateName)
	fileInit := createAndAppendObject(fileNameInit)

	// Create a new LoggerSync instance
	logger := &LoggerSync{
		tag:              tag,
		enDebug:          debugMode,
		writeFilesEnable: false,
		gateName:         gateName,
		filesName:        fileNameInit,
		file:             fileInit,
	}
	log.SetOutput(os.Stdout)
	log.SetFlags(0) // Disable the default timestamp and log prefix
	return logger
}

func (l *LoggerSync) ChangeFileRoutine() {
	HOUR := 00
	MINUTE := 00
	go func() {
		for range time.Tick(1 * time.Minute) {
			hours, minutes, _ := time.Now().Clock()
			if hours == HOUR && minutes == MINUTE {
				// Close first previous object file
				l.file.Close()

				// Create new file object with the append mode
				l.filesName = fileNameGenerator(l.gateName)
				l.file = createAndAppendObject(l.filesName)
			}
		}
	}()

}

func (l *LoggerSync) writeLog(msg string) {
	if l.writeFilesEnable {
		l.file.WriteString(msg + "\n")
	}
}

func (l *LoggerSync) SetWriteFilesEnable(enable bool) {
	l.writeFilesEnable = enable
}

func (l *LoggerSync) applyStyle(str string, styles ...int8) string {
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

func (l *LoggerSync) getTime() string {
	return "[" + time.Now().Format("2006-01-02 15:04:05.000") + "] "
}

func (l *LoggerSync) SetInfoStyle(styles ...int8) {
	for i := range len(styles) {
		l.infoStyle = append(l.infoStyle, styles[i])
	}
}

func (l *LoggerSync) SetWarnStyle(styles ...int8) {
	for i := range len(styles) {
		l.warnStyle = append(l.warnStyle, styles[i])
	}
}

func (l *LoggerSync) SetErrorStyle(styles ...int8) {
	for i := range len(styles) {
		l.errorStyle = append(l.errorStyle, styles[i])
	}
}

func (l *LoggerSync) SetDebugStyle(styles ...int8) {
	for i := range len(styles) {
		l.debugStyle = append(l.debugStyle, styles[i])
	}
}

func (l *LoggerSync) SetPanicStyle(styles ...int8) {
	for i := range len(styles) {
		l.panicStyle = append(l.panicStyle, styles[i])
	}
}

func (l *LoggerSync) SetFatalStyle(styles ...int8) {
	for i := range len(styles) {
		l.fatalStyle = append(l.fatalStyle, styles[i])
	}
}

func (l *LoggerSync) SetDefaultStyle() {
	l.SetInfoStyle(StyleFgWhite)
	l.SetWarnStyle(StyleFgYellow)
	l.SetErrorStyle(StyleFgRed)
	l.SetDebugStyle(StyleFontItalic, StyleFontComment)
	l.SetPanicStyle(StyleFontBold, StyleFgBlack, StyleBgMagenta)
	l.SetFatalStyle(StyleFontBold, StyleFgBlack, StyleBgRed)
}

// LOG FORMAT
// [INFO] [TIME] [TAG]: message

func (l *LoggerSync) Info(a ...any) {
	msg := l.getTime() + infoKey + l.tag + ": " + fmt.Sprint(a...)
	l.writeLog(msg)
	msg = l.applyStyle(msg, l.infoStyle...)
	log.Println(msg)
}

func (l *LoggerSync) Infof(format string, a ...any) {
	msg := l.getTime() + infoKey + l.tag + ": " + fmt.Sprintf(format, a...)
	l.writeLog(msg)
	msg = l.applyStyle(msg, l.infoStyle...)
	log.Println(msg)
}

func (l *LoggerSync) Warn(a ...any) {
	msg := l.getTime() + warnKey + l.tag + ": " + fmt.Sprint(a...)
	l.writeLog(msg)
	msg = l.applyStyle(msg, l.warnStyle...)
	log.Println(msg)
}

func (l *LoggerSync) Warnf(format string, a ...any) {
	msg := l.getTime() + warnKey + l.tag + ": " + fmt.Sprintf(format, a...)
	l.writeLog(msg)
	msg = l.applyStyle(msg, l.warnStyle...)
	log.Println(msg)
}

func (l *LoggerSync) Error(a ...any) {
	msg := l.getTime() + errorKey + l.tag + ": " + fmt.Sprint(a...)
	l.writeLog(msg)
	msg = l.applyStyle(msg, l.errorStyle...)
	log.Println(msg)
}

func (l *LoggerSync) Errorf(format string, a ...any) {
	msg := l.getTime() + errorKey + l.tag + ": " + fmt.Sprintf(format, a...)
	l.writeLog(msg)
	msg = l.applyStyle(msg, l.errorStyle...)
	log.Println(msg)
}

func (l *LoggerSync) Debug(a ...any) {
	if l.enDebug {
		msg := l.getTime() + debugKey + l.tag + ": " + fmt.Sprint(a...)
		l.writeLog(msg)
		msg = l.applyStyle(msg, l.debugStyle...)
		log.Println(msg)
	}
}

func (l *LoggerSync) Debugf(format string, a ...any) {
	if l.enDebug {
		msg := l.getTime() + debugKey + l.tag + ": " + fmt.Sprintf(format, a...)
		l.writeLog(msg)
		msg = l.applyStyle(msg, l.debugStyle...)
		log.Println(msg)
	}
}

func (l *LoggerSync) Panic(a ...any) {
	msg := l.getTime() + panicKey + l.tag + ": " + fmt.Sprint(a...)
	l.writeLog(msg)
	msg = l.applyStyle(msg, l.panicStyle...)
	log.Panicln(msg)
}

func (l *LoggerSync) Panicf(format string, a ...any) {
	msg := l.getTime() + panicKey + l.tag + ": " + fmt.Sprintf(format, a...)
	l.writeLog(msg)
	msg = l.applyStyle(msg, l.panicStyle...)
	log.Panicln(msg)
}

func (l *LoggerSync) Fatal(a ...any) {
	msg := l.getTime() + fatalKey + l.tag + ": " + fmt.Sprint(a...)
	l.writeLog(msg)
	msg = l.applyStyle(msg, l.fatalStyle...)
	log.Fatalln(msg)
}

func (l *LoggerSync) Fatalf(format string, a ...any) {
	msg := l.getTime() + fatalKey + l.tag + ": " + fmt.Sprintf(format, a...)
	l.writeLog(msg)
	msg = l.applyStyle(msg, l.fatalStyle...)
	log.Fatalln(msg)
}
