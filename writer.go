package log

import (
	"taylz.io/log/writer"
	"taylz.io/types"
)

// NewRollingWriter creates a io.WriteCloser that rotates backing file, named for each day, for a directory
func NewRollingWriter(path string) types.WriteCloser {
	return writer.NewRoller(path)
}
