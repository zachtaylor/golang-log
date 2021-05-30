package log

import "taylz.io/types"

// Step is a log builder
type Step struct {
	finish StepFinisher
	fields types.Dict
}

// StepFinisher is a hook for Step
type StepFinisher interface {
	// Finish completes a line
	Finish(Level, types.Dict, types.Source, []interface{})
}

// NewStep creates a new Step
func NewStep(finisher StepFinisher, fields types.Dict) Step {
	if fields == nil {
		fields = make(types.Dict)
	}
	return Step{
		finish: finisher,
		fields: fields,
	}
}

// New copies the Step
func (log Step) New() types.Logger { return log.Copy() }

// Copy returns a shallow copy of this Step
func (log Step) Copy() Step {
	copy := NewStep(log.finish, nil)
	for k, v := range log.fields {
		copy.fields[k] = v
	}
	return copy
}

// Add writes any value to the fields
func (log Step) Add(k string, v interface{}) types.Logger {
	copy := log.Copy()
	copy.fields[k] = v
	return copy
}

// With writes all values to the fields
func (log Step) With(fields types.Dict) types.Logger {
	copy := log.Copy()
	for k, v := range fields {
		copy.fields[k] = v
	}
	return copy
}

// Trace writes a log with LevelTrace
func (log Step) Trace(args ...interface{}) { log.help(LevelTrace, args) }

// Debug writes a log with LevelDebug
func (log Step) Debug(args ...interface{}) { log.help(LevelDebug, args) }

// Info writes a log with LevelInfo
func (log Step) Info(args ...interface{}) { log.help(LevelInfo, args) }

// Warn writes a log with LevelWarn
func (log Step) Warn(args ...interface{}) { log.help(LevelWarn, args) }

// Error writes a log with LevelError
func (log Step) Error(args ...interface{}) { log.help(LevelError, args) }

// Out writes a log with LevelOut
func (log Step) Out(args ...interface{}) { log.help(LevelOut, args) }

func (log Step) help(level Level, args []interface{}) {
	log.finish.Finish(level, log.fields, types.NewSource(2), args)
}
