package log

import (
	"fmt"
	"sort"
	"strings"
)

// ColorFormat is a Formatter designed for humans
type ColorFormat struct {
	Colors     ColorMap
	ColorTime  bool
	ColorSrc   bool
	ColorMsg   bool
	ColorField bool
	SrcFmt     SourceFormatter
	TimeFmt    TimeFormatter
}

// DefaultColorFormat builds a ColorFormat with some default options
func DefaultColorFormat(sourcer SourceFormatter, timer TimeFormatter) Formatter {
	return &ColorFormat{
		Colors:     DefaultColorMap(),
		ColorMsg:   true,
		ColorField: true,
		SrcFmt:     sourcer,
		TimeFmt:    timer,
	}
}

func (cf *ColorFormat) Format(line Line) []byte {
	useColor := cf.Colors != nil
	var colorCode, time, src string
	if useColor {
		if code := cf.Colors[line.Level]; code == "" {
			useColor = false
		} else {
			colorCode = code
		}
	}
	if cf.TimeFmt != nil {
		time = cf.TimeFmt.FormatTime(line.Time)
	}
	if cf.SrcFmt != nil {
		src = cf.SrcFmt.FormatSource(line.Source)
	}

	var sb strings.Builder
	var dirty bool // internal: buffer wants spacing at current write index

	if !useColor {
		if len(time) > 0 {
			sb.WriteString(time)
			sb.WriteByte(' ')
		}
		sb.WriteByte(line.Level.ByteCode())
		dirty = true // unsafe sb.WriteByte(' ')
		if len(src) > 0 {
			sb.WriteByte(' ')
			sb.WriteString(src)
		}
		for _, arg := range line.Args {
			sb.WriteByte(' ')
			cf.writeValue(&sb, useColor, arg)
		}
	} else {
		var incolor bool

		if len(time) > 0 {
			if cf.ColorTime {
				sb.WriteString(colorCode)
				incolor = true
			}
			sb.WriteString(time)
			dirty = true
		}

		if len(src) > 0 {
			if incolor && !cf.ColorSrc {
				sb.WriteString(ColorOff)
				incolor = false
			}
			if dirty {
				sb.WriteByte(' ')
			}
			if cf.ColorSrc && !incolor {
				sb.WriteString(colorCode)
				incolor = true
			}
			sb.WriteString(src)
			dirty = true // unsafe sb.WriteByte(' ')
		}

		if incolor && !cf.ColorMsg {
			sb.WriteString(ColorOff)
			incolor = false
		}
		if dirty {
			sb.WriteByte(' ')
			dirty = false
		}
		if cf.ColorMsg && !incolor {
			sb.WriteString(colorCode)
			incolor = true
		}

		for _, arg := range line.Args {
			if dirty {
				sb.WriteByte(' ')
			}
			cf.writeValue(&sb, useColor, arg)
			dirty = true
		}

		if incolor {
			sb.WriteString(ColorOff)
			incolor = false
		}
	}

	if dirty {
		sb.WriteString("  \t\t")
		dirty = false
	}

	for _, key := range cf.fieldKeys(line.Fields) {
		if dirty {
			sb.WriteByte(' ')
		}
		if useColor && cf.ColorField {
			sb.WriteString(colorCode)
		}
		sb.WriteString(key)
		if useColor {
			sb.WriteString(ColorGray)
			sb.WriteByte('=')
			sb.WriteString(ColorOff)
		}
		cf.writeValue(&sb, useColor, line.Fields[key])
		dirty = true
	}

	sb.WriteByte('\n')

	return []byte(sb.String())
}

func (*ColorFormat) fieldKeys(fields Fields) []string {
	keys, i := make([]string, len(fields)), 0
	for k := range fields {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

// writeValue records a value representing any type
func (*ColorFormat) writeValue(sb *strings.Builder, useColor bool, arg interface{}) {
	if arg == nil {
		if useColor {
			sb.WriteString(TextBold)
		}
		sb.WriteString("<nil>")
		if useColor {
			sb.WriteString(ColorOff)
		}
	} else if err, _ := arg.(error); err != nil {
		if useColor {
			sb.WriteString(ColorRed)
		}
		sb.WriteString(err.Error())
		if useColor {
			sb.WriteString(ColorOff)
		}
	} else if str, _ := arg.(string); len(str) > 0 {
		sb.WriteString(str)
	} else if stringer, _ := arg.(fmt.Stringer); stringer != nil {
		sb.WriteString(stringer.String())
	} else {
		fmt.Fprintf(sb, "%v", arg)
	}
}
