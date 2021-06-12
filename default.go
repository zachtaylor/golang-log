package log

import (
	"io"
	"strings"
)

// Default creates a basic logger
func Default() Writer {
	return Lining(StdOutLiner(DefaultColorFormat(ClearSourceFormatter(), ClearTimeFormatter())))
}

// Lining creates *T with default clock
func Lining(liner Liner) *T {
	return New(DefaultClock(), liner)
}

// DefaultColorFormat builds a ColorFormat with some default options
func DefaultColorFormat(sourcer SourceFormatter, timer TimeFormatter) Formatter {
	return &ColorFormat{
		Colors:     DefaultColorMap(),
		ColorMsg:   true,
		ColorField: true,
		SrcFmt:     sourcer,
		TimeFmt:    timer,
	}
}

// Classic uses the Lining and ClassicLiner
func Classic(lvl Level, gopath string, w io.WriteCloser) *T {
	return Lining(ClassicLiner(lvl, gopath, w))
}

// ClassicLiner uses LevelLiner, IOLiner, DefaultColorFormat, ClassicSourceFormatter, and DefaultTimeFormatter
func ClassicLiner(lvl Level, gopath string, w io.WriteCloser) Liner {
	return LevelLiner(lvl, IOLiner(DefaultColorFormat(ClassicSourceFormatter(gopath), DefaultTimeFormatter()), w))
}

const classicSrcLen = 24

// ClassicSourceFormatter uses detail format, exact length, and path prefix removal
func ClassicSourceFormatter(gopaths ...string) SourceFormatter {
	return RestringSourceFormatter(
		DetailSourceFormatter(),
		RestringerMiddleware(RestringerCutPrefixes(gopaths), RestringerLenExact(classicSrcLen)),
	)
}

// SourceFormatterLen adds RestringerLenExact as a middleware layer above SourceFormatter
func SourceFormatterLen(srcfmt SourceFormatter, len int) SourceFormatter {
	return RestringSourceFormatter(srcfmt, RestringerLenExact(len))
}

// DefaultModuleRoot grabs the callstack and makes assumptions about directory structure
//
// call from "{ROOT}/cmd/example/main.go"
func DefaultModuleRoot() string {
	filePath := NewSource(1).File()
	const parentno = 2
	for i := 0; strings.Contains(filePath, "/") && i <= parentno; i++ {
		filePath = filePath[:strings.LastIndex(filePath, "/")]
	}
	return filePath
}
