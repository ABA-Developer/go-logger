// package logger provide async non blocking logging for better app performance
package logger

const (
	infoKey  = "[INFO ] "
	warnKey  = "[WARN ] "
	errorKey = "[ERROR] "
	debugKey = "[DEBUG] "
	panicKey = "[PANIC] "
	fatalKey = "[FATAL] "
)

const (
	StyleFontNone      int8 = 0
	StyleFontBold      int8 = 1
	StyleFontComment   int8 = 2
	StyleFontItalic    int8 = 3
	StyleFontUnderline int8 = 4
	StyleFgBlack       int8 = 30
	StyleFgRed         int8 = 31
	StyleFgGreen       int8 = 32
	StyleFgYellow      int8 = 33
	StyleFgBlue        int8 = 34
	StyleFgMagenta     int8 = 35
	StyleFgCyan        int8 = 36
	StyleFgWhite       int8 = 37
	StyleFgDefault     int8 = 39
	StyleBgBlack       int8 = 40
	StyleBgRed         int8 = 41
	StyleBgGreen       int8 = 42
	StyleBgYellow      int8 = 43
	StyleBgBlue        int8 = 44
	StyleBgMagenta     int8 = 45
	StyleBgCyan        int8 = 46
	StyleBgWhite       int8 = 47
	StyleBgDefault     int8 = 49
)
