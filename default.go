package log

import (
	"strings"

	"taylz.io/log/restringer"
	"taylz.io/types"
)

// Default creates a basic logger
func Default() *T {
	return Lining(StdOutLiner(DefaultColorFormat(ClearSourceFormatter(), ClearTimeFormatter())))
}

// Lining creates *T with default clock
func Lining(liner Liner) *T {
	return New(types.DefaultClock(), liner)
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
func Classic(lvl Level, gopath string, w types.Writer) *T {
	return Lining(ClassicLiner(lvl, gopath, w))
}

// ClassicLiner uses LevelLiner, IOLiner, DefaultColorFormat, ClassicSourceFormatter, and DefaultTimeFormatter
func ClassicLiner(lvl Level, gopath string, w types.Writer) Liner {
	return LevelLiner(lvl, IOLiner(DefaultColorFormat(ClassicSourceFormatter(gopath), DefaultTimeFormatter()), w))
}

const classicSrcLen = 24

// ClassicSourceFormatter uses detail format, exact length, and path prefix removal
func ClassicSourceFormatter(gopaths ...string) SourceFormatter {
	return RestringSourceFormatter(
		DetailSourceFormatter(),
		restringer.Middleware(restringer.CutPrefixList(gopaths), restringer.LenExact(classicSrcLen)),
	)
}

// SourceFormatterLen adds restringer.LenExact as a middleware layer above SourceFormatter
func SourceFormatterLen(srcfmt SourceFormatter, len int) SourceFormatter {
	return RestringSourceFormatter(srcfmt, restringer.LenExact(len))
}

// DefaultModuleRoot grabs the callstack and makes assumptions about directory structure
//
// call from "{ROOT}/cmd/example/main.go"
func DefaultModuleRoot() string {
	filePath := types.NewSource(1).File()
	const parentno = 2
	for i := 0; strings.Contains(filePath, "/") && i <= parentno; i++ {
		filePath = filePath[:strings.LastIndex(filePath, "/")]
	}
	return filePath
}
