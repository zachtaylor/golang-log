package log

import (
	"runtime"
	"strconv"
)

// Source links to a go code instruction
type Source struct {
	file string
	line int
}

// NewSource creates a Source reference to the caller number # higher in the stack frame than the invoking call (0 is default)
func NewSource(history int) Source {
	_, file, line, _ := runtime.Caller(1 + history)
	return Source{
		file: file,
		line: line,
	}
}

// File returns go file reference
func (src Source) File() string { return src.file }

// Line returns source line
func (src Source) Line() int { return src.line }

func (src Source) String() string { return src.file + "#" + strconv.FormatInt(int64(src.line), 10) }
