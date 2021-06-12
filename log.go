package log

// T is a customizable log
type T struct {
	clock Clock
	liner Liner
}

// New creates a Log
func New(clock Clock, liner Liner) *T { return &T{clock: clock, liner: liner} }

// New returns a new Logger
func (t *T) New() Writer { return NewStep(t, nil) }

// Add returns a new Logger with a K/V pair
func (t *T) Add(k string, v interface{}) Writer { return NewStep(t, Fields{k: v}) }

// With writes all values to the Fields
func (t *T) With(fields Fields) Writer { return NewStep(t, fields) }

// Trace writes a log with LevelTrace
func (t *T) Trace(args ...interface{}) { t.Finish(LevelTrace, Fields{}, NewSource(1), args) }

// Debug writes a log with LevelDebug
func (t *T) Debug(args ...interface{}) { t.Finish(LevelDebug, Fields{}, NewSource(1), args) }

// Info writes a log with LevelInfo
func (t *T) Info(args ...interface{}) { t.Finish(LevelInfo, Fields{}, NewSource(1), args) }

// Warn writes a log with LevelWarn
func (t *T) Warn(args ...interface{}) { t.Finish(LevelWarn, Fields{}, NewSource(1), args) }

// Error writes a log with LevelError
func (t *T) Error(args ...interface{}) { t.Finish(LevelError, Fields{}, NewSource(1), args) }

// Out writes a log with LevelOut
func (t *T) Out(args ...interface{}) { t.Finish(LevelOut, Fields{}, NewSource(1), args) }

// Finish implements line.Finisher
func (t *T) Finish(lvl Level, fs Fields, src Source, args []interface{}) {
	t.liner.Line(NewLine(lvl, fs, src, t.clock.Now(), args))
}

// Close calls the Liner.Close
func (t *T) Close() error { return t.liner.Close() }
