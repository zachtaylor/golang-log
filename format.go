package log

import (
	"strconv"

	"taylz.io/types"
)

// Formatter is used to format log lines
type Formatter interface {
	Format(Line) []byte
}

// SourceFormatter is a strategy for converting types.Source to string
type SourceFormatter interface {
	FormatSource(types.Source) string
}

// SourceFormatterString is a string replacement for SourceFormatter
type SourceFormatterString string

// Format implements SourceFormatter
func (str SourceFormatterString) FormatSource(types.Source) string { return string(str) }

// SourceFormatterFunc is a func type implementing SourceFormatter
type SourceFormatterFunc func(types.Source) string

// Format implements SourceFormatter by calling the function
func (f SourceFormatterFunc) FormatSource(src types.Source) string { return f(src) }

// RestringSourceFormatter attaches types.Restringer to SourceFormatter
func RestringSourceFormatter(srcfmt SourceFormatter, restringer types.Restringer) SourceFormatter {
	return SourceFormatterFunc(func(src types.Source) string { return restringer.Restring(srcfmt.FormatSource(src)) })
}

// ClearSourceFormatter returns a SourceFormatter that only produces empty string
func ClearSourceFormatter() SourceFormatter {
	return SourceFormatterString("")
}

// SimpleSourceFormatter returns a SourceFormatter that returns src.String()
func SimpleSourceFormatter() SourceFormatter {
	return SourceFormatterFunc(func(src types.Source) string { return src.String() })
}

// PrettySourceFormatter returns a SourceFormatter that returns src.File() with ".go" suffix removed, if present
func PrettySourceFormatter() SourceFormatter {
	return SourceFormatterFunc(func(src types.Source) string {
		f := src.File()
		if lenf := len(f); lenf < 5 {
			return src.String()
		} else if f[lenf-3:] != ".go" {
			return src.String()
		} else {
			return f[:lenf-3]
		}
	})
}

// DetailSourceFormatter returns a SourceFormatter that returns src.File() [without ".go"] + ":" + src.Line()
func DetailSourceFormatter() SourceFormatter {
	return SourceFormatterFunc(func(src types.Source) string {
		f := src.File()
		if lenf := len(f); lenf < 5 {
			return src.String()
		} else if f[lenf-3:] != ".go" {
			return src.String()
		} else {
			return f[:lenf-3] + ":" + strconv.FormatInt(int64(src.Line()), 10)
		}
	})
}

// TimeFormatter is used to format time values
type TimeFormatter interface {
	FormatTime(types.Time) string
}

// TimeFormatterFunc is a func type that implements TimeFormatter
type TimeFormatterFunc func(t types.Time) string

// Format implements TimeFormatter using the func ptr
func (f TimeFormatterFunc) FormatTime(t types.Time) string { return f(t) }

// TimeFormatterConstant is used to return constant string as a TimeFormatter
type TimeFormatterConstant string

// Format implements TimeFormatter by returning the same string
func (str TimeFormatterConstant) FormatTime(t types.Time) string { return string(str) }

// TimeFormatString is a string that passes itself to time.Format
type TimeFormatString string

// Format returns time package formatting given this format string
func (str TimeFormatString) FormatTime(time types.Time) string {
	return time.Format(string(str))
}

// ClearTimeFormatter returns a TimeFormatter that only produces empty string
func ClearTimeFormatter() TimeFormatter {
	return TimeFormatterConstant("")
}

// DefaultTimeFormatter returns a TimeFormatter that uses using 24-hour time format ("15:04:05")
func DefaultTimeFormatter() TimeFormatter {
	return TimeFormatString("15:04:05")
}
