package log

import "os"

// Terminal is short for With(LevelLiner(lvl, TerminalLiner()))
func Terminal(lvl Level) Writer {
	return With(LevelLiner(lvl, TerminalLiner()))
}

// TerminalLiner uses LevelLiner, IOLiner, DefaultColorFormat, ClassicSourceFormatter, and DefaultTimeFormatter
func TerminalLiner() Liner {
	return IOLiner(
		DefaultColorFormat(
			TwentyFiveSourceFormatter(),
			DefaultTimeFormatter(),
		), os.Stdout)
}
