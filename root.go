package log

// Root is a customizable log
type Root struct {
	clock Clock
	liner Liner
}

// NewRoot creates a Root
func NewRoot(clock Clock, liner Liner) *Root { return &Root{clock: clock, liner: liner} }

// New returns a new Writer
func (t *Root) New() Writer { return NewBranch(t, nil) }

// Add returns a new Writer with a K/V pair
func (t *Root) Add(k string, v interface{}) Writer { return NewBranch(t, Fields{k: v}) }

// With writes all values to the Fields
func (t *Root) With(fields Fields) Writer { return NewBranch(t, fields) }

func (t *Root) Prefix(args ...interface{}) Writer { return NewBranch(t, Fields{}, args...) }

// Trace writes a log with LevelTrace
func (t *Root) Trace(args ...interface{}) { t.Finish(LevelTrace, Fields{}, NewSource(1), args) }

// Debug writes a log with LevelDebug
func (t *Root) Debug(args ...interface{}) { t.Finish(LevelDebug, Fields{}, NewSource(1), args) }

// Info writes a log with LevelInfo
func (t *Root) Info(args ...interface{}) { t.Finish(LevelInfo, Fields{}, NewSource(1), args) }

// Warn writes a log with LevelWarn
func (t *Root) Warn(args ...interface{}) { t.Finish(LevelWarn, Fields{}, NewSource(1), args) }

// Error writes a log with LevelError
func (t *Root) Error(args ...interface{}) { t.Finish(LevelError, Fields{}, NewSource(1), args) }

// Out writes a log with LevelOut
func (t *Root) Out(args ...interface{}) { t.Finish(LevelOut, Fields{}, NewSource(1), args) }

// Finish implements line.Finisher
func (t *Root) Finish(lvl Level, fs Fields, src Source, args []interface{}) {
	t.liner.Line(NewLine(lvl, fs, src, t.clock.Now(), args))
}

// Close calls the Liner.Close
func (t *Root) Close() error { return t.liner.Close() }
