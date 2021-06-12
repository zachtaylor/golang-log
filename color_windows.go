// +build windows

package log

import (
	"syscall"

	sequences "github.com/nine-lives-later/go-windows-terminal-sequences"
)

func init() { sequences.EnableVirtualTerminalProcessing(syscall.Stdout, true) }
