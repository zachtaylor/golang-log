// +build windows

package log

import (
	"syscall"

	sequences "github.com/konsorten/go-windows-terminal-sequences"
)

func init() { sequences.EnableVirtualTerminalProcessing(syscall.Stdout, true) }
