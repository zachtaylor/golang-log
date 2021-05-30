package log

const (
	// ColorOff is a color coding literal
	ColorOff = "\x1b[0m"
	// ColorBlack is a color coding literal
	ColorBlack = "\x1b[30m"
	// ColorRed is a color coding literal
	ColorRed = "\x1b[31m"
	// ColorGreen is a color coding literal
	ColorGreen = "\x1b[32m"
	// ColorYellow is a color coding literal
	ColorYellow = "\x1b[33m"
	// ColorBlue is a color coding literal
	ColorBlue = "\x1b[34m"
	// ColorMagenta is a color coding literal
	ColorMagenta = "\x1b[35m"
	// ColorCyan is a color coding literal
	ColorCyan = "\x1b[36m"
	// ColorLightGray is a color coding literal
	ColorLightGray = "\x1b[37m"
	// ColorGray is a color coding literal
	ColorGray = "\x1b[90m"
	// ColorLightRed is a color coding literal
	ColorLightRed = "\x1b[91m"
	// ColorLightGreen is a color coding literal
	ColorLightGreen = "\x1b[92m"
	// ColorLightYellow is a color coding literal
	ColorLightYellow = "\x1b[93m"
	// ColorLightBlue is a color coding literal
	ColorLightBlue = "\x1b[94m"
	// ColorLightMagenta is a color coding literal
	ColorLightMagenta = "\x1b[95m"
	// ColorLightCyan is a color coding literal
	ColorLightCyan = "\x1b[96m"
	// ColorWhite is a color coding literal
	ColorWhite = "\x1b[97m"
)

// ColorMap is coloration guide
type ColorMap = map[Level]string

// DefaultColorMap returns the default color set
func DefaultColorMap() ColorMap {
	return ColorMap{
		LevelTrace: ColorLightMagenta,
		LevelDebug: ColorLightBlue,
		LevelInfo:  ColorLightGreen,
		LevelWarn:  ColorLightYellow,
		LevelError: ColorLightRed,
		LevelOut:   ColorWhite,
	}
}
