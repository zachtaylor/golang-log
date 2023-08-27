package log

type Writer interface {
	New() Writer
	Add(string, interface{}) Writer
	With(Fields) Writer
	Prefix(...interface{}) Writer
	Trace(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Out(...interface{})
	Close() error
}
