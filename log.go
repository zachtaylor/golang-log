package log

// Default creates a basic logger for standard out
func Default() Writer {
	return With(StdOutLiner(DefaultColorFormat(PrettySourceFormatter(), DefaultTimeFormatter())))
}

// With creates *T with default clock
func With(liner Liner) Writer {
	return NewRoot(DefaultClock(), liner)
}
