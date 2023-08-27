package log

// Level is used to rank the importance of logs
type Level int8

const (
	// LevelTrace is the lowest level
	LevelTrace Level = iota
	// LevelDebug is a more detailed value
	LevelDebug
	// LevelInfo is the default level
	LevelInfo
	// LevelError is a raised level
	LevelWarn
	// LevelError is the considered the top level
	LevelError
	// LevelOut is the highest value, a sentinal value
	LevelOut
)

// ByteCode returns an ASCII byte code for this level
func (t Level) ByteCode() byte {
	switch t {
	case LevelTrace:
		return 84 // T
	case LevelDebug:
		return 68 // D
	case LevelInfo:
		return 73 // I
	case LevelWarn:
		return 87 // W
	case LevelError:
		return 69 // E
	case LevelOut:
		return 79 // O
	default:
		return 63 // ?
	}
}

// Lowercase returns a lower case string for
func (t Level) Lowercase() string {
	switch t {
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelOut:
		return "out"
	default:
		return "<unknown>"
	}
}

// GetLevel returns the level named, if valid
//
// valid values: "t", "T", "trace", "TRACE", etc...
func GetLevel(string string) (Level, error) {
	switch string {
	case "t":
		fallthrough
	case "T":
		fallthrough
	case "trace":
		fallthrough
	case "TRACE":
		return LevelTrace, nil
	case "d":
		fallthrough
	case "D":
		fallthrough
	case "debug":
		fallthrough
	case "DEBUG":
		return LevelDebug, nil
	case "i":
		fallthrough
	case "I":
		fallthrough
	case "info":
		fallthrough
	case "INFO":
		return LevelInfo, nil
	case "w":
		fallthrough
	case "W":
		fallthrough
	case "warn":
		fallthrough
	case "WARN":
		return LevelWarn, nil
	case "e":
		fallthrough
	case "E":
		fallthrough
	case "error":
		fallthrough
	case "ERROR":
		return LevelError, nil
	case "o":
		fallthrough
	case "O":
		fallthrough
	case "out":
		fallthrough
	case "OUT":
		return LevelOut, nil
	default:
		return -1, ErrLevelUnknown(string)
	}
}

// ErrUnknown is an error type returned by GetLevel
type ErrLevelUnknown string

func (val ErrLevelUnknown) Error() string { return "level unknown: " + string(val) }
