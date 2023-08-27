package log

import "errors"

// Branch is a non-root stage in creating a Line
type Branch struct {
	finish Finisher
	fields Fields
	prefix []interface{}
}

// Finisher is a hook for creating a Line
type Finisher interface {
	// Finish completes a line
	Finish(Level, Fields, Source, []interface{})
}

// NewBranch creates a new Step
func NewBranch(finisher Finisher, fields Fields, prefixArgs ...interface{}) Branch {
	if fields == nil {
		fields = Fields{}
	}
	return Branch{
		finish: finisher,
		fields: fields,
		prefix: prefixArgs,
	}
}

// New copies the Step
func (log Branch) New() Writer { return log.Copy() }

// Copy returns a shallow copy of this Step
func (log Branch) Copy() Branch {
	copy := Branch{
		finish: log.finish,
		fields: Fields{},
		prefix: make([]interface{}, len(log.prefix)),
	}
	for k, v := range log.fields {
		copy.fields[k] = v
	}
	for i, v := range log.prefix {
		copy.prefix[i] = v
	}
	return copy
}

// Add writes any value to the fields
func (log Branch) Add(k string, v interface{}) Writer {
	copy := log.Copy()
	copy.fields[k] = v
	return copy
}

// With writes all values to the fields
func (log Branch) With(fields Fields) Writer {
	copy := log.Copy()
	for k, v := range fields {
		copy.fields[k] = v
	}
	return copy
}

func (log Branch) Prefix(args ...interface{}) Writer {
	copy := log.Copy()
	copy.prefix = append(copy.prefix, args...)
	return copy
}

// Trace writes a log with LevelTrace
func (log Branch) Trace(args ...interface{}) { log.help(LevelTrace, args) }

// Debug writes a log with LevelDebug
func (log Branch) Debug(args ...interface{}) { log.help(LevelDebug, args) }

// Info writes a log with LevelInfo
func (log Branch) Info(args ...interface{}) { log.help(LevelInfo, args) }

// Warn writes a log with LevelWarn
func (log Branch) Warn(args ...interface{}) { log.help(LevelWarn, args) }

// Error writes a log with LevelError
func (log Branch) Error(args ...interface{}) { log.help(LevelError, args) }

// Out writes a log with LevelOut
func (log Branch) Out(args ...interface{}) { log.help(LevelOut, args) }

func (log Branch) help(level Level, args []interface{}) {
	log.finish.Finish(level, log.fields, NewSource(2), append(log.prefix, args...))
}

// ErrCannotCloseStep is returned by Step.Close
var ErrCannotCloseStep = errors.New("cannot close log step")

func (log Branch) Close() error { return ErrCannotCloseStep }
