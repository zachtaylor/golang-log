package log

// Dual is short for With(LevelLiner(lvl, DualLiner(path)))
func Dual(lvl Level, path string) Writer {
	return With(LevelLiner(lvl, DualLiner(path)))
}

// DualLiner logs both Terminal and JsonRotatingFile
func DualLiner(path string) Liner {
	return MultiLiner{
		TerminalLiner(),
		JsonRotatingFileLiner(nil, path),
	}
}
