package log

import (
	"strconv"
	"time"
)

// Formatter is used to format log lines
type Formatter interface {
	Format(Line) []byte
}

// FormatterFunc is a func type implementing Formatter
type FormatterFunc func(Line) []byte

// Format implements Formatter by calling the function
func (f FormatterFunc) Format(l Line) []byte { return f(l) }

// SourceFormatter is a strategy for converting Source to string
type SourceFormatter interface {
	FormatSource(Source) string
}

// SourceFormatterFunc is a func type implementing SourceFormatter
type SourceFormatterFunc func(Source) string

// Format implements SourceFormatter by calling the function
func (f SourceFormatterFunc) FormatSource(src Source) string { return f(src) }

// RestringSourceFormatter attaches Restringer to SourceFormatter
func RestringSourceFormatter(srcfmt SourceFormatter, restringer Restringer) SourceFormatter {
	return SourceFormatterFunc(func(src Source) string {
		return restringer.Restring(srcfmt.FormatSource(src))
	})
}

// SimpleSourceFormatter returns a SourceFormatter that returns src.String()
func SimpleSourceFormatter() SourceFormatter {
	return SourceFormatterFunc(func(src Source) string { return src.String() })
}

// PrettySourceFormatter returns a SourceFormatter that doesn't show line numbers
func PrettySourceFormatter() SourceFormatter {
	return SourceFormatterFunc(func(src Source) string {
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

// DefaultSourceFormatter returns a SourceFormatter that removes ".go" from the
// file name (if present), saving line number
func DefaultSourceFormatter() SourceFormatter {
	return SourceFormatterFunc(func(src Source) string {
		f := src.File()
		if lenf := len(f); lenf < 5 {
			return src.String()
		} else if f[lenf-3:] != ".go" {
			return src.String()
		} else {
			return f[:lenf-3] + "#" + strconv.FormatInt(int64(src.Line()), 10)
		}
	})
}

// TwentyFiveSourceFormatter uses DefaultSourceFormatter with exact length 25
func TwentyFiveSourceFormatter() SourceFormatter {
	return RestringSourceFormatter(
		DefaultSourceFormatter(),
		RestringerLenExact(25),
	)
}

// TimeFormatter is used to format time values
type TimeFormatter interface {
	FormatTime(time.Time) string
}

// TimeFormatterFunc is a func type that implements TimeFormatter
type TimeFormatterFunc func(t time.Time) string

// Format implements TimeFormatter using the func ptr
func (f TimeFormatterFunc) FormatTime(t time.Time) string { return f(t) }

// TimeFormatString is a string that passes itself to time.Format
type TimeFormatString string

// Format returns time package formatting given this format string
func (str TimeFormatString) FormatTime(time time.Time) string { return time.Format(string(str)) }

// DefaultTimeFormatter returns a TimeFormatter that uses using 24-hour time format ("15:04:05")
func DefaultTimeFormatter() TimeFormatter { return TimeFormatString("15:04:05") }

// FormatStringLenExact returns a string of set size, elided (from the left) if longer, or right-padded if shorter
func FormatStringLenExact(str string, size int) string {
	lenstr := len(str)
	if lenstr == size {
		return str
	} else if size < 1 {
		return ""
	}
	lendif := lenstr - size
	buf := make([]byte, size)
	var i, j int
	if lendif > 0 {
		buf[0], buf[1], buf[2] = '.', '.', '.'
		i = 3
		j = lendif + i
	}
	for i < size && j < lenstr {
		buf[i] = str[j]
		i++
		j++
	}
	for ; i < size; i++ {
		buf[i] = ' '
	}
	return string(buf)
}

// FormatStringLenMin returns strings of minimum size, right-padded if shorter
func FormatStringLenMin(str string, size int) string {
	lenstr := len(str)
	if lenstr >= size {
		return str
	} else if size < 1 {
		return ""
	}
	buf := make([]byte, size)
	for i := 0; i < lenstr; i++ {
		buf[i] = str[i]
	}
	for i := lenstr; i < size; i++ {
		buf[i] = ' '
	}
	return string(buf)
}
