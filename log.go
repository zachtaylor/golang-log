package log

import "taylz.io/types"

// T is a customizable log
type T struct {
	clock types.Clocker
	liner Liner
}

// New creates a Log
func New(clock types.Clocker, liner Liner) *T { return &T{clock: clock, liner: liner} }

// New returns a new Logger
func (t *T) New() types.Logger { return NewStep(t, nil) }

// Add returns a new Logger with a K/V pair
func (t *T) Add(k string, v interface{}) types.Logger { return NewStep(t, Fields{k: v}) }

// With writes all values to the Fields
func (t *T) With(fields Fields) types.Logger { return NewStep(t, fields) }

// Trace writes a log with LevelTrace
func (t *T) Trace(args ...interface{}) { t.Finish(LevelTrace, Fields{}, types.NewSource(1), args) }

// Debug writes a log with LevelDebug
func (t *T) Debug(args ...interface{}) { t.Finish(LevelDebug, Fields{}, types.NewSource(1), args) }

// Info writes a log with LevelInfo
func (t *T) Info(args ...interface{}) { t.Finish(LevelInfo, Fields{}, types.NewSource(1), args) }

// Warn writes a log with LevelWarn
func (t *T) Warn(args ...interface{}) { t.Finish(LevelWarn, Fields{}, types.NewSource(1), args) }

// Error writes a log with LevelError
func (t *T) Error(args ...interface{}) { t.Finish(LevelError, Fields{}, types.NewSource(1), args) }

// Out writes a log with LevelOut
func (t *T) Out(args ...interface{}) { t.Finish(LevelOut, Fields{}, types.NewSource(1), args) }

// Finish implements line.Finisher
func (t *T) Finish(lvl Level, fs Fields, src types.Source, args []interface{}) {
	t.liner.Line(NewLine(lvl, fs, src, t.clock.NewTime(), args))
}
